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
	iAdd      // +
	iDec      // -
	iMul      // *
	iDiv      // ÷
	iMod      // %
	iSal      // <<
	iSar      // >>
	iBitAnd   // &
	iBitAndOr // &^
	iBitXor   // ^
	iBitOr    // |
	iBt       // >
	iLt       // <

	iEq // ==

	iOr  // ||
	iAnd // &&

	// 19 actions
)

type byteCode struct {
	ins     action
	operand interface{}
}

type zVal interface{}

type virtualMachine struct {
	instructions []byteCode
	pc           int
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

var actionRuner = [21]actionHandler{
	iPop:      zPop,
	iNot:      zNot,
	iAdd:      zAdd,
	iDec:      zDec,
	iMul:      zMul,
	iDiv:      zDiv,
	iMod:      zMod,
	iSal:      zSal,
	iSar:      zSar,
	iBitAnd:   zBitAnd,
	iBitAndOr: nil,
	iBitXor:   zBitXor,
	iBitOr:    zBitOr,
	iEq:       zEq,
	iBt:       zBt,
	iLt:       zLt,
	iOr:       zOr,
	iAnd:      zAnd,
}

func zBitAnd(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val, ok := vm.stack[top].(int)
	if !ok {
		return errors.New("operand of & must be an int")
	}
	dst, ok := vm.stack[top-1].(int)
	if !ok {
		return errors.New("operand of & must be an int")
	}
	vm.stack[top-1] = dst & val
	vm.stack = vm.stack[:top]
	return nil
}

func zBitAndOr(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val, ok := vm.stack[top].(int)
	if !ok {
		return errors.New("operand of | must be an int")
	}
	dst, ok := vm.stack[top-1].(int)
	if !ok {
		return errors.New("operand of | must be an int")
	}
	vm.stack[top-1] = dst &^ val
	vm.stack = vm.stack[:top]
	return nil
}

func zBitOr(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val, ok := vm.stack[top].(int)
	if !ok {
		return errors.New("operand of | must be an int")
	}
	dst, ok := vm.stack[top-1].(int)
	if !ok {
		return errors.New("operand of | must be an int")
	}
	vm.stack[top-1] = dst | val
	vm.stack = vm.stack[:top]
	return nil
}

func zBitXor(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val, ok := vm.stack[top].(int)
	if !ok {
		return errors.New("operand of ^ must be an int")
	}
	dst, ok := vm.stack[top-1].(int)
	if !ok {
		return errors.New("operand of ^ must be an int")
	}
	vm.stack[top-1] = dst ^ val
	vm.stack = vm.stack[:top]
	return nil
}

func zSal(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val, ok := vm.stack[top].(int)
	if !ok {
		return errors.New("operand of << must be an int")
	}
	dst, ok := vm.stack[top-1].(int)
	if !ok {
		return errors.New("operand of << must be an int")
	}
	vm.stack[top-1] = dst << uint(val)
	vm.stack = vm.stack[:top]
	return nil
}

func zSar(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	val, ok := vm.stack[top].(int)
	if !ok {
		return errors.New("operand of << must be an int")
	}
	dst, ok := vm.stack[top-1].(int)
	if !ok {
		return errors.New("operand of << must be an int")
	}
	vm.stack[top-1] = dst >> uint(val)
	vm.stack = vm.stack[:top]
	return nil
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
	val, ok = env[fields[0]]
	if !ok {
		return nil, false
	}
	nextVal := reflect.ValueOf(val)
	for nextVal.Type().Kind() == reflect.Ptr {
		nextVal = nextVal.Elem()
	}
	for _, fieldName := range fields[1:] {
		nextVal = nextVal.FieldByName(fieldName)
		if !nextVal.IsValid() {
			ok = false
			return
		}
	}
	val = nextVal.Interface()
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
	switch ra := vm.stack[top].(type) {
	case int:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = ra + rb
		case float64:
			vm.stack[top-1] = float64(ra) + rb
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
	case float64:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = ra + float64(rb)
		case float64:
			vm.stack[top-1] = ra + rb
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
	default:
		return errOperandOfNumericOperationIsNotIntOrFloat64
	}
	vm.stack = vm.stack[:top]
	return nil
}

func zDec(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	switch ra := vm.stack[top].(type) {
	case int:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = rb - ra
		case float64:
			vm.stack[top-1] = rb - float64(ra)
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
	case float64:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = float64(rb) - ra
		case float64:
			vm.stack[top-1] = rb - ra
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
	default:
		return errOperandOfNumericOperationIsNotIntOrFloat64
	}
	vm.stack = vm.stack[:top]
	return nil
}

func zMul(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	switch ra := vm.stack[top].(type) {
	case int:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = ra * rb
		case float64:
			vm.stack[top-1] = float64(ra) * rb
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
	case float64:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = ra * float64(rb)
		case float64:
			vm.stack[top-1] = ra * rb
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
	default:
		return errOperandOfNumericOperationIsNotIntOrFloat64
	}
	vm.stack = vm.stack[:top]
	return nil
}

func zDiv(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	switch ra := vm.stack[top].(type) {
	case int:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = rb / ra
		case float64:
			vm.stack[top-1] = rb / float64(ra)
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
	case float64:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = float64(rb) / ra
		case float64:
			vm.stack[top-1] = rb / ra
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
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
	switch ra := vm.stack[top].(type) {
	case int:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = rb > ra
		case float64:
			vm.stack[top-1] = rb > float64(ra)
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
	case float64:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = float64(rb) > ra
		case float64:
			vm.stack[top-1] = rb > ra
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
	default:
		return errOperandOfNumericOperationIsNotIntOrFloat64
	}
	vm.stack = vm.stack[:top]
	return nil
}

func zLt(vm *virtualMachine) error {
	top := len(vm.stack) - 1
	switch ra := vm.stack[top].(type) {
	case int:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = rb < ra
		case float64:
			vm.stack[top-1] = rb < float64(ra)
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
	case float64:
		switch rb := vm.stack[top-1].(type) {
		case int:
			vm.stack[top-1] = float64(rb) < ra
		case float64:
			vm.stack[top-1] = rb < ra
		default:
			return errOperandOfNumericOperationIsNotIntOrFloat64
		}
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

func (l *byteCodeListener) ExitFirstExpr(ctx *grammar.FirstExprContext) {
	code := byteCode{}
	op := ctx.GetOp().GetText()
	switch op {
	case "*":
		code.ins = iMul
	case "/":
		code.ins = iDiv
	case "%":
		code.ins = iMod
	case "<<":
		code.ins = iSal
	case ">>":
		code.ins = iSar
	case "&":
		code.ins = iBitAnd
	case "&^":
		code.ins = iBitAndOr
	}
	l.stack = append(l.stack, code)
}

func (l *byteCodeListener) ExitSecondExpr(ctx *grammar.SecondExprContext) {
	code := byteCode{}
	op := ctx.GetOp().GetText()
	switch op {
	case "+":
		code.ins = iAdd
	case "-":
		code.ins = iDec
	case "|":
		code.ins = iBitOr
	case "^":
		code.ins = iBitXor

	}
	l.stack = append(l.stack, code)
}

func (l *byteCodeListener) ExitThirdExpr(ctx *grammar.ThirdExprContext) {
	op := ctx.GetOp().GetText()
	switch op {
	case "==":
		l.stack = append(l.stack, byteCode{ins: iEq})
	case "!=":
		l.stack = append(l.stack, byteCode{ins: iEq}, byteCode{ins: iNot})
	case "<":
		l.stack = append(l.stack, byteCode{ins: iLt})
	case "<=":
		l.stack = append(l.stack, byteCode{ins: iBt}, byteCode{ins: iNot})
	case ">":
		l.stack = append(l.stack, byteCode{ins: iBt})
	case ">=":
		l.stack = append(l.stack, byteCode{ins: iLt}, byteCode{ins: iNot})
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
