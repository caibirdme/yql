package lambda

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	grammar "github.com/caibirdme/yql/internal/lambda"
)

type action uint8

const (

	// single operand
	_ action = iota
	iPush
	vPush
	iPop
	iNot // !

	// binary operand
	iAdd // +
	iDec // -
	iMul // *
	iDiv // ÷
	iMod // %
	iSal // <<
	iSar // >>
	iBt  // >
	iLt  // <

	iEq // ==

	iOr  // ||
	iAnd // &&

	// 13 actions
)

type byteCode struct {
	ins     action
	operand interface{}
}

type zVal interface{}

type virtualMachine struct {
	instructions []byteCode
	pc           int
	overflow     bool
	err          error
	stack        []zVal
	env          map[string]interface{}
	params       []string
}

func (vm *virtualMachine) mainLoop() error {
	length := len(vm.instructions)
	var err error
	vm.pc = 0
	for vm.pc < length {
		instruction := vm.instructions[vm.pc]
		err = vm.exec(instruction)
		if nil != err {
			return err
		}
		vm.pc++
	}
	return nil
}

type variableType struct {
	varName string
}

func (vm *virtualMachine) exec(instruction byteCode) error {
	if instruction.ins == iPush {
		return zPush(vm, instruction.operand)
	} else if instruction.ins == vPush {
		varName, ok := instruction.operand.(string)
		if !ok {
			return errRuntimeErr
		}
		return variablePush(vm, varName)
	}
	return actionRuner[instruction.ins](vm)
}

type actionHandler func(*virtualMachine) error

var actionRuner = [17]actionHandler{
	iPop: zPop,
	iNot: zNot,
	iAdd: zAdd,
	iDec: zDec,
	iMul: zMul,
	iDiv: zDiv,
	iMod: zMod,
	iSal: nil,
	iSar: nil,
	iEq:  zEq,
	iBt:  zBt,
	iLt:  zLt,
	iOr:  zOr,
	iAnd: zAnd,
}

func zPush(vm *virtualMachine, operand zVal) error {
	vm.stack = append(vm.stack, operand)
	return nil
}

func variablePush(vm *virtualMachine, varName string) error {
	val, ok := resolveVar(vm.env, varName)
	if !ok {
		return fmt.Errorf("%s not found", varName)
	}
	vm.stack = append(vm.stack, val)
	return nil
}

func resolveVar(env map[string]interface{}, varName string) (val interface{}, ok bool) {
	fields := strings.Split(varName, ".")
	val, ok = env[varName]
	if !ok {
		return nil, false
	}
	nextVal := reflect.ValueOf(val)
	for _, fieldName := range fields[1:] {
		nextVal = nextVal.FieldByName(fieldName)
		if !nextVal.IsValid() {
			ok = false
			return
		}
	}
	t, ok := val.(int64)
	if ok {
		return int(t), true
	}
	return val, true
}

func zPop(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	vm.stack = vm.stack[:top]
	return nil
}

func zNot(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val := vm.stack[top]
	switch reflect.TypeOf(val).Kind() {
	case reflect.Bool:
		p := val.(bool)
		vm.stack[top] = !p
		return nil
	}
	return errOperandOfNotIsNotBool
}

func zAdd(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val := vm.stack[top]
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		p1 := val.(int)
		p2 := vm.stack[top-1].(int)
		vm.stack[top-1] = p1 + p2
	case reflect.Float64:
		p1 := val.(float64)
		p2 := vm.stack[top-1].(float64)
		vm.stack[top-1] = p1 + p2
	default:
		return errOperandOfNumericOperationIsNotIntOrFloat64
	}
	vm.stack = vm.stack[:top]
	return nil
}

func zDec(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val := vm.stack[top]
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		p1 := val.(int)
		p2 := vm.stack[top-1].(int)
		vm.stack[top-1] = p2 - p1
	case reflect.Float64:
		p1 := val.(float64)
		p2 := vm.stack[top-1].(float64)
		vm.stack[top-1] = p2 - p1
	default:
		return errOperandOfNumericOperationIsNotIntOrFloat64
	}
	vm.stack = vm.stack[:top]
	return nil
}

func zMul(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val := vm.stack[top]
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		p1 := val.(int)
		p2 := vm.stack[top-1].(int)
		vm.stack[top-1] = p1 * p2
	case reflect.Float64:
		p1 := val.(float64)
		p2 := vm.stack[top-1].(float64)
		vm.stack[top-1] = p1 * p2
	default:
		return errOperandOfNumericOperationIsNotIntOrFloat64
	}
	vm.stack = vm.stack[:top]
	return nil
}

func zDiv(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val := vm.stack[top]
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		p1 := val.(int)
		p2 := vm.stack[top-1].(int)
		vm.stack[top-1] = p2 / p1
	case reflect.Float64:
		p1 := val.(float64)
		p2 := vm.stack[top-1].(float64)
		vm.stack[top-1] = p2 / p1
	default:
		return errOperandOfNumericOperationIsNotIntOrFloat64
	}
	vm.stack = vm.stack[:top]
	return nil
}

func zEq(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	vm.stack[top-1] = vm.stack[top-1] == vm.stack[top]
	vm.stack = vm.stack[:top]
	return nil
}

func zBt(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val := vm.stack[top]
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		p1 := val.(int)
		p2 := vm.stack[top-1].(int)
		vm.stack[top-1] = p2 > p1
	case reflect.Float64:
		p1 := val.(float64)
		p2 := vm.stack[top-1].(float64)
		vm.stack[top-1] = p2 > p1
	default:
		return errOperandOfNumericOperationIsNotIntOrFloat64
	}
	vm.stack = vm.stack[:top]
	return nil
}

func zLt(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val := vm.stack[top]
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		p1 := val.(int)
		p2 := vm.stack[top-1].(int)
		vm.stack[top-1] = p2 < p1
	case reflect.Float64:
		p1 := val.(float64)
		p2 := vm.stack[top-1].(float64)
		vm.stack[top-1] = p2 < p1
	default:
		return errOperandOfNumericOperationIsNotIntOrFloat64
	}
	vm.stack = vm.stack[:top]
	return nil
}

func zOr(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val := vm.stack[top]
	switch reflect.TypeOf(val).Kind() {
	case reflect.Bool:
		p1 := val.(bool)
		p2 := vm.stack[top-1].(bool)
		vm.stack[top-1] = p1 || p2
		vm.stack = vm.stack[:top]
		return nil
	}
	return errOperandOfOrIsNotBool
}

func zAnd(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val := vm.stack[top]
	switch reflect.TypeOf(val).Kind() {
	case reflect.Bool:
		p1 := val.(bool)
		p2 := vm.stack[top-1].(bool)
		vm.stack[top-1] = p1 && p2
		vm.stack = vm.stack[:top]
		return nil
	}
	return errOperandOfAndIsNotBool
}

func zMod(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val := vm.stack[top]
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		p1 := val.(int)
		p2 := vm.stack[top-1].(int)
		vm.stack[top-1] = p2 % p1
	default:
		return errOperandOfModIsNotInt
	}
	vm.stack = vm.stack[:top]
	return nil
}

var (
	errOperandOfNotIsNotBool                      = errors.New("Operand Of Not IsNot Bool")
	errOperandOfAndIsNotBool                      = errors.New("Operand Of And IsNot Bool")
	errOperandOfOrIsNotBool                       = errors.New("Operand Of Or IsNot Bool")
	errOperandOfNumericOperationIsNotIntOrFloat64 = errors.New("Operand Of Numeric Operation IsNot Int Or Float64")
	errOperandOfModIsNotInt                       = errors.New("Operand Of Mod IsNot Int")
)

type byteCodeListener struct {
	*grammar.BaseLambdaListener
	stack  []byteCode
	params []string
}

func (l *byteCodeListener) ExitSignature(ctx *grammar.SignatureContext) {
	text := ctx.GetText()
	if text[0] == '(' {
		text = text[1:]
	}
	t := len(text) - 1
	if text[t] == ')' {
		text = text[:t]
	}
	for _, s := range strings.Split(text, ",") {
		l.params = append(l.params, strings.Trim(s, " "))
	}
}

func (l *byteCodeListener) ExitOrExpr(ctx *grammar.OrExprContext) {
	l.stack = append(l.stack, byteCode{
		ins: iOr,
	})
}

func (l *byteCodeListener) ExitAndExpr(ctx *grammar.AndExprContext) {
	l.stack = append(l.stack, byteCode{
		ins: iAnd,
	})
}

func (l *byteCodeListener) ExitShiftExpr(ctx *grammar.ShiftExprContext) {
	text := ctx.GetOp().GetText()
	var code byteCode
	switch text {
	case "<<":
		code.ins = iSal
	case ">>":
		code.ins = iSar
	}
	l.stack = append(l.stack, code)
}

func (l *byteCodeListener) ExitEqExpre(ctx *grammar.EqExpreContext) {
	l.stack = append(l.stack, byteCode{
		ins: iEq,
	})
}

func (l *byteCodeListener) ExitMulExpr(ctx *grammar.MulExprContext) {
	text := ctx.GetOp().GetText()
	var code byteCode
	switch text {
	case "*":
		code.ins = iMul
	case "/":
		code.ins = iDiv
	case "%":
		code.ins = iMod
	}
	l.stack = append(l.stack, code)
}

func (l *byteCodeListener) ExitAddExpr(ctx *grammar.AddExprContext) {
	text := ctx.GetOp().GetText()
	var code byteCode
	switch text {
	case "+":
		code.ins = iAdd
	case "-":
		code.ins = iDec
	}
	l.stack = append(l.stack, code)
}

func (l *byteCodeListener) ExitCompareExpr(ctx *grammar.CompareExprContext) {
	text := ctx.GetOp().GetText()
	switch text {
	case ">":
		l.stack = append(l.stack, byteCode{ins: iBt})
	case ">=":
		l.stack = append(l.stack, byteCode{ins: iLt})
		l.stack = append(l.stack, byteCode{ins: iNot})
	case "<":
		l.stack = append(l.stack, byteCode{ins: iLt})
	case "<=":
		l.stack = append(l.stack, byteCode{ins: iBt})
		l.stack = append(l.stack, byteCode{ins: iNot})
	}
}

func (l *byteCodeListener) ExitValueExpr(ctx *grammar.ValueExprContext) {
	code := byteCode{
		ins: iPush,
	}
	text := ctx.GetText()
	switch {
	case isIntLit(text):
		t, _ := strconv.ParseInt(text, 10, 64)
		code.operand = int(t)
	case isFloatLit(text):
		t, _ := strconv.ParseFloat(text, 64)
		code.operand = t
	case isStringLit(text):
		if text[0] == '`' {
			code.operand = strings.Trim(text, "`")
		} else {
			code.operand = strings.Trim(text, `"`)
		}
	case isBoolLit(text):
		t, _ := strconv.ParseBool(text)
		code.operand = t
	default:
		l.stack = append(l.stack, byteCode{
			ins:     vPush,
			operand: text,
		})
		return
	}
	l.stack = append(l.stack, code)
}

var (
	reINT    = regexp.MustCompile(`^\d+$`)
	reFLOAT  = regexp.MustCompile(`^\d+\.?\d*$`)
	reSTRING = regexp.MustCompile("^(`[^`]*`|\".*\")$")
)

func isIntLit(text string) bool {
	return reINT.Match([]byte(text))
}

func isFloatLit(text string) bool {
	return reFLOAT.Match([]byte(text))
}

func isStringLit(text string) bool {
	return reSTRING.Match([]byte(text))
}

func isBoolLit(text string) bool {
	return text == "true" || text == "false"
}
