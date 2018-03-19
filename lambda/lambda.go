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
	codeSegment, params, err := sdt(root)
	if nil != err {
		return newErrMState(err)
	}
	return &MState{
		code:   codeSegment,
		params: params,
		cb:     filterRunner,
	}
}

func filterRunner(vm *virtualMachine, params ...interface{}) Result {
	l := len(params)
	if 0 == l {
		return errZeroParams
	}
	targetType := reflect.TypeOf(params[0])
	if targetType.Kind() != reflect.Slice {
		return errTargetNotSlice
	}
	targetData := reflect.ValueOf(params[0])
	length := targetData.Len()
	if 0 == length {
		return errEmptySlice
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
			return &result{err: err}
		}
		if len(vm.stack) < 1 {
			return &result{err: errRuntimeErr}
		}
		val := vm.stack[0]
		if reflect.TypeOf(val).Kind() != reflect.Bool {
			return &result{err: errResultTypeErr}
		}
		b := val.(bool)
		if b {
			res = reflect.Append(res, targetData.Index(i))
		}
	}
	return &result{
		data: res.Interface(),
		err:  nil,
	}
}

var (
	errRuntimeErr    = errors.New("runtime error")
	errResultTypeErr = errors.New("result type error")
)

// Syntax-Directed Translation
func sdt(root antlr.Tree) (code []byteCode, params []string, err error) {
	walker := &byteCodeListener{
		stack: make([]byteCode, 0, 5),
	}
	defer func() {
		if r := recover(); nil != r {
			err = fmt.Errorf("%+v", r)
		}
	}()
	antlr.ParseTreeWalkerDefault.Walk(walker, root)
	return walker.stack, walker.params, nil
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
	root, err := getTreeRoot(funcStmt)
	if nil != err {
		return newErrMState(err)
	}
	_, err = typeCheck(root)
	if nil != err {
		return newErrMState(err)
	}
	codeSegment, params, err := sdt(root)
	if nil != err {
		return newErrMState(err)
	}
	return &MState{
		code:   codeSegment,
		params: params,
		cb:     mapRunner,
	}
}

var (
	errZeroParams     = &result{err: errors.New("params shouldn't be empty")}
	errTargetNotSlice = &result{err: errors.New("the first params should be a slice")}
	errEmptySlice     = &result{err: errors.New("slice is empty")}
)

func mapRunner(vm *virtualMachine, params ...interface{}) Result {
	l := len(params)
	if 0 == l {
		return errZeroParams
	}
	targetType := reflect.TypeOf(params[0])
	if targetType.Kind() != reflect.Slice {
		return errTargetNotSlice
	}
	targetData := reflect.ValueOf(params[0])
	length := targetData.Len()
	if 0 == length {
		return errEmptySlice
	}
	elementType := targetType.Elem()
	if elementType.Kind() == reflect.Struct {
		return nil
	}
	res := reflect.MakeSlice(reflect.SliceOf(elementType), 0, length)
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
			return &result{err: err}
		}
		if len(vm.stack) < 1 {
			return &result{data: nil, err: errRuntimeErr}
		}
		val := vm.stack[0]
		if reflect.TypeOf(val) != elementType {
			return &result{data: nil, err: errResultTypeErr}
		}
		res = reflect.Append(res, reflect.ValueOf(val))
	}
	return &result{
		data: res.Interface(),
		err:  nil,
	}
}

type Result interface {
	Interface() (interface{}, error)
}

// Result store the result for a call
type result struct {
	data interface{}
	err  error
}

func (r *result) Interface() (interface{}, error) {
	return r.data, r.err
}

// MState is a middle state
// It records the call stack and generates responding bytecode
type MState struct {
	code   []byteCode
	params []string
	err    error
	cb     func(*virtualMachine, ...interface{}) Result
	next   *MState
}

func (m *MState) Filter(funcStmt string) *MState {
	if m.err != nil {
		return m
	}
	t := m
	for t.next != nil {
		t = t.next
	}
	newMState := Filter(funcStmt)
	if newMState.err != nil {
		return newMState
	}
	t.next = newMState
	return m
}

func (m *MState) Map(funcStmt string) *MState {
	if m.err != nil {
		return m
	}
	t := m
	for t.next != nil {
		t = t.next
	}
	newMState := Map(funcStmt)
	if newMState.err != nil {
		return newMState
	}
	t.next = newMState
	return m
}

// Call creates the context and exec the function
func (m *MState) Call(params ...interface{}) Result {
	if m.err != nil {
		return &result{data: nil, err: m.err}
	}
	t := m
	var r Result
	for t != nil {
		vm := &virtualMachine{
			instructions: t.code,
			params:       t.params,
		}
		r = t.cb(vm, params...)
		newR, err := r.Interface()
		if nil != err {
			return r
		}
		params[0] = newR
		t = t.next
	}
	return r
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

func (l *typeCheckListener) ExitFisrtExpr(ctx *grammar.FirstExprContext) {
	top := len(l.stack) - 1
	t := l.stack[top] & l.stack[top-1]
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	var availableType uint
	availableType = 1 << reflect.Int
	op := ctx.GetOp().GetText()
	switch op {
	case "*", "/":
		availableType |= 1 << reflect.Float64
	}
	t &= availableType
	if t == 0 {
		l.halt = true
		panic(true)
	}
	l.stack[top-1] = t
	l.stack = l.stack[:top]
}

func (l *typeCheckListener) ExitSecondExpr(ctx *grammar.SecondExprContext) {
	top := len(l.stack) - 1
	t := l.stack[top] & l.stack[top-1]
	if t == 0 {
		l.halt = true
		panic(l.halt)
	}
	var availableType uint
	availableType = 1 << reflect.Int
	op := ctx.GetOp().GetText()
	switch op {
	case "+", "-":
		availableType |= 1 << reflect.Float64
	}
	t &= availableType
	if t == 0 {
		l.halt = true
		panic(true)
	}
	l.stack[top-1] = t
	l.stack = l.stack[:top]
}

func (l *typeCheckListener) ExitThirdExpr(ctx *grammar.ThirdExprContext) {
	top := len(l.stack) - 1
	t := l.stack[top] & l.stack[top-1]
	if t == 0 {
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
		t = 1<<reflect.Int | 1<<reflect.Float64
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
	1<<reflect.Float64 |
	1<<reflect.String)
