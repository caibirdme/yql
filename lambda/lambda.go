package lambda

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	grammar "github.com/caibirdme/yql/internal/lambda"
)

type bailLexer struct {
	*grammar.LambdaLexer
}

// Override default Recover and do nothing
// lexer can quit when any error occur
func (l *bailLexer) Recover(r antlr.RecognitionException) {
	panic(r)
}

var (
	bailErrStrategy = antlr.NewBailErrorStrategy()
)

// Filter creates a new slice with all elements that
// pass the test implemented by the provided function
func Filter(funcStmt string) *MState {
	root, err := getTreeRoot(funcStmt)
	if nil != err {
		return newErrMState(err)
	}
	returnType, err := typeCheck(root)
	if nil != err {
		return newErrMState(err)
	}
	if returnType&(1<<reflect.Bool) == 0 {
		return newErrMState(errors.New("Error Return Type"))
	}
	vm, err := sdt(root)
	if nil != err {
		return newErrMState(err)
	}
	return &MState{
		vm: vm,
		cb: filterRunner,
	}
}

func filterRunner(vm *virtualMachine, params ...interface{}) *Result {
	l := len(params)
	if 0 == l {
		return nil
	}
	targetType := reflect.TypeOf(params[0])
	if targetType.Kind() != reflect.Slice {
		return nil
	}
	targetData := reflect.ValueOf(params[0])
	length := targetData.Len()
	if 0 == length {
		return nil
	}
	res := reflect.MakeSlice(reflect.SliceOf(targetType.Elem()), 0, length>>1)
	for i := 0; i < length; i++ {
		vm.env = map[string]interface{}{
			vm.params[0]: targetData.Index(i).Interface(),
		}
		for j := 1; j < len(vm.params); j++ {
			vm.env[vm.params[j]] = params[j]
		}
		vm.stack = vm.stack[:0]
		err := vm.mainLoop()
		if nil != err {
			return nil
		}
		if len(vm.stack) < 1 {
			return &Result{data: nil, err: errRuntimeErr}
		}
		val := vm.stack[0]
		if reflect.TypeOf(val).Kind() != reflect.Bool {
			return &Result{data: nil, err: errResultTypeErr}
		}
		b := val.(bool)
		if b {
			res = reflect.Append(res, targetData.Index(i))
		}
	}
	return &Result{
		data: res.Interface(),
		err:  nil,
	}
}

var (
	errRuntimeErr    = errors.New("runtime error")
	errResultTypeErr = errors.New("result type error")
)

// Syntax-Directed Translation
func sdt(root antlr.Tree) (vm *virtualMachine, err error) {
	walker := &byteCodeListener{
		stack: make([]byteCode, 0, 18),
	}
	defer func() {
		if r := recover(); nil != r {
			err = fmt.Errorf("%+v", r)
		}
	}()
	antlr.ParseTreeWalkerDefault.Walk(walker, root)
	return &virtualMachine{
		instructions: walker.stack,
		params:       walker.params,
	}, nil
}

func getTreeRoot(funcStmt string) (tree grammar.ILambdaContext, err error) {
	defer func() {
		if r := recover(); nil != r {
			err = fmt.Errorf("%+v", r)
		}
	}()
	inputStream := antlr.NewInputStream(funcStmt)
	lexer := &bailLexer{grammar.NewLambdaLexer(inputStream)}
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewLambdaParser(stream)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(bailErrStrategy)
	tree = parser.Lambda()
	return tree, err
}

// Map creates a new slice with the results of calling
// a provided function on every element in the calling slice
func Map(funcStmt string) *MState {
	return nil
}

// Result store the result for a call
type Result struct {
	data interface{}
	err  error
}

func (r *Result) Interface() (interface{}, error) {
	return r.data, r.err
}

// MState is a middle state
// It records the call stack and generates responding bytecode
type MState struct {
	vm  *virtualMachine
	err error
	cb  func(*virtualMachine, ...interface{}) *Result
}

func (m *MState) Call(params ...interface{}) *Result {
	if m.err != nil {
		return &Result{data: nil, err: m.err}
	}
	return m.cb(m.vm, params...)
}

func newErrMState(err error) *MState {
	return &MState{
		err: err,
	}
}

func typeCheck(root antlr.Tree) (acceptableType uint, err error) {
	walker := &typeCheckListener{
		stack: make([]uint, 0, 4),
	}
	defer func() {
		if r := recover(); nil != r {
			err = fmt.Errorf("%s %+v", errTypeMisMatch, r)
		}
	}()
	antlr.ParseTreeWalkerDefault.Walk(walker, root)
	if len(walker.stack) > 0 {
		return walker.stack[0], nil
	}
	return 0, errors.New("stack calc error")
}

var (
	errTypeMisMatch = errors.New("type mismatch")
)

type typeCheckListener struct {
	*grammar.BaseLambdaListener
	stack []uint
	halt  bool
}

func (l *typeCheckListener) ExitOrExpr(ctx *grammar.OrExprContext) {
	top := len(l.stack) - 1
	t := l.stack[top] & l.stack[top-1]
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	t &= 1 << reflect.Bool
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	l.stack[top-1] = t
	l.stack = l.stack[:top]
}

func (l *typeCheckListener) ExitAndExpr(ctx *grammar.AndExprContext) {
	top := len(l.stack) - 1
	t := l.stack[top] & l.stack[top-1]
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	t &= 1 << reflect.Bool
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	l.stack[top-1] = t
	l.stack = l.stack[:top]
}

func (l *typeCheckListener) ExitShiftExpr(ctx *grammar.ShiftExprContext) {
	top := len(l.stack) - 1
	t := l.stack[top] & l.stack[top-1]
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	t &= (1<<reflect.Int | 1<<reflect.Int64)
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	l.stack[top-1] = t
	l.stack = l.stack[:top]
}

func (l *typeCheckListener) ExitEqExpre(ctx *grammar.EqExpreContext) {
	top := len(l.stack) - 1
	t := l.stack[top] & l.stack[top-1]
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	l.stack[top-1] = 1 << reflect.Bool
	l.stack = l.stack[:top]
}

func (l *typeCheckListener) ExitMulExpr(ctx *grammar.MulExprContext) {
	top := len(l.stack) - 1
	t := l.stack[top] & l.stack[top-1]
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	t &= (1<<reflect.Int | 1<<reflect.Int64 | 1<<reflect.Float64)
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	l.stack[top-1] = t
	l.stack = l.stack[:top]
}

func (l *typeCheckListener) ExitAddExpr(ctx *grammar.AddExprContext) {
	top := len(l.stack) - 1
	t := l.stack[top] & l.stack[top-1]
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	t &= (1<<reflect.Int | 1<<reflect.Int64 | 1<<reflect.Float64)
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	l.stack[top-1] = t
	l.stack = l.stack[:top]
}

func (l *typeCheckListener) ExitCompareExpr(ctx *grammar.CompareExprContext) {
	top := len(l.stack) - 1
	if l.stack[top]&l.stack[top-1] == 0 {
		l.halt = true
		panic(l.halt)
	}
	l.stack[top-1] = 1 << reflect.Bool
	l.stack = l.stack[:top]
}

func (l *typeCheckListener) ExitBasicLit(ctx *grammar.BasicLitContext) {
	var t uint
	switch {
	case ctx.INT_LIT() != nil:
		t = 1<<reflect.Int | 1<<reflect.Int64
	case ctx.FLOAT_LIT() != nil:
		t = 1 << reflect.Float64
	case ctx.STRING_LIT() != nil:
		t = 1 << reflect.String
	case ctx.BOOL_LIT() != nil:
		t = 1 << reflect.Bool
	}
	l.stack = append(l.stack, t)
}

func (l *typeCheckListener) ExitOperandName(ctx *grammar.OperandNameContext) {
	l.stack = append(l.stack, anyType)
}

var anyType uint = (1<<reflect.Bool |
	1<<reflect.Int |
	1<<reflect.Int64 |
	1<<reflect.Float64 |
	1<<reflect.String)
