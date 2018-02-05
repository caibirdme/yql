package yql

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/caibirdme/yql/internal/grammar"
	"github.com/caibirdme/yql/internal/stack"
)

type boolStack interface {
	Push(bool)
	Pop() bool
	// Init initialize the stack, considering the stack will be reused
	Init()
}

type yqlListener struct {
	*grammar.BaseYqlListener
	stack       boolStack
	data        map[string]interface{}
	notFoundErr error
}

func (l *yqlListener) ExitBooleanExpr(ctx *grammar.BooleanExprContext) {
	if l.notFoundErr != nil {
		return
	}
	operator := ctx.GetOp().GetText()
	fieldName := ctx.FIELDNAME().GetText()
	actualValue, ok := l.data[fieldName]
	if !ok {
		l.notFoundErr = fmt.Errorf("%s not provided", fieldName)
		return
	}
	allValue := ctx.AllValue()
	valueArr := make([]string, 0, len(allValue))
	for _, v := range allValue {
		valueArr = append(valueArr, v.GetText())
	}
	res := compare(actualValue, valueArr, operator)
	l.stack.Push(res)
}

func (l *yqlListener) ExitOrExpr(ctx *grammar.OrExprContext) {
	if l.notFoundErr != nil {
		return
	}
	q1 := l.stack.Pop()
	q2 := l.stack.Pop()
	l.stack.Push(q1 || q2)
}

func (l *yqlListener) ExitAndExpr(ctx *grammar.AndExprContext) {
	if l.notFoundErr != nil {
		return
	}
	q1 := l.stack.Pop()
	q2 := l.stack.Pop()
	l.stack.Push(q1 && q2)
}

type lexerErrHandler struct {
	*antlr.DefaultErrorListener
}

func (d *lexerErrHandler) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(e)
}

func match(rawYQL string, data map[string]interface{}) (ok bool, err error) {
	if 0 == len(data) {
		return false, nil
	}
	defer func() {
		if r := recover(); nil != r {
			ok = false
			err = fmt.Errorf("%v", r)
		}
	}()
	inputStream := antlr.NewInputStream(rawYQL)
	lexer := grammar.NewYqlLexer(inputStream)
	lexer.AddErrorListener(&lexerErrHandler{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	})
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewYqlParser(stream)
	parser.BuildParseTrees = true
	parser.SetErrorHandler(antlr.NewBailErrorStrategy())
	tree := parser.Query()
	l := &yqlListener{stack: stack.NewStack()}
	l.stack.Init()
	l.data = data
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	if l.notFoundErr != nil {
		return false, l.notFoundErr
	}
	return l.stack.Pop(), nil
}

// Match interprete the rawYQL and execute it with the provided data
// error will be returned if rawYQL contains illegal syntax
func Match(rawYQL string, data map[string]interface{}) (bool, error) {
	return match(rawYQL, data)
}

func compare(actualValue interface{}, expectValue []string, op string) bool {
	if len(expectValue) > 1 {
		return compareSet(actualValue, expectValue, op)
	}
	e := removeQuote(expectValue[0])
	switch actual := actualValue.(type) {
	case int:
		expect, err := strconv.ParseInt(e, 10, 64)
		if nil != err {
			return false
		}
		return cmpInt(int64(actual), expect, op)
	case int64:
		expect, err := strconv.ParseInt(e, 10, 64)
		if nil != err {
			return false
		}
		return cmpInt(actual, expect, op)
	case float64:
		expect, err := strconv.ParseFloat(e, 64)
		if nil != err {
			return false
		}
		return cmpFloat(actual, expect, op)
	case string:
		return cmpStr(actual, e, op)
	default:
		return false
	}
}

func removeQuote(str string) string {
	l := 0
	r := len(str)
	if str[0] == '\'' {
		l++
	}
	if str[r-1] == '\'' {
		r--
	}
	return str[l:r]
}
