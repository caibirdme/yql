package yql

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/caibirdme/yql/grammar"
	"github.com/caibirdme/yql/stack"
)

type yqlListener struct {
	*grammar.BaseYqlListener
	stack stack.Stack
	data  map[string]interface{}
	halt  bool
}

func (l *yqlListener) ExitBooleanExpr(ctx *grammar.BooleanExprContext) {
	if l.halt {
		return
	}
	operator := ctx.GetOp().GetText()
	fieldName := ctx.FIELDNAME().GetText()
	actualValue, ok := l.data[fieldName]
	if !ok {
		l.halt = true
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
	if l.halt {
		return
	}
	q1 := l.stack.Pop()
	q2 := l.stack.Pop()
	l.stack.Push(q1 || q2)
}

func (l *yqlListener) ExitAndExpr(ctx *grammar.AndExprContext) {
	if l.halt {
		return
	}
	q1 := l.stack.Pop()
	q2 := l.stack.Pop()
	l.stack.Push(q1 && q2)
}

func match(rawYQL string, data map[string]interface{}) bool {
	if 0 == len(data) {
		return false
	}
	inputStream := antlr.NewInputStream(rawYQL)
	lexer := grammar.NewYqlLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewYqlParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Query()
	l := &yqlListener{stack: stack.NewStack()}
	l.stack.Init()
	l.data = data
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	if l.halt {
		return false
	}
	return l.stack.Pop()
}

func compare(actualValue interface{}, expectValue []string, op string) bool {
	if len(expectValue) > 1 {
		return compareSet(actualValue, expectValue, op)
	}
	e := removeQuote(expectValue[0])
	switch actual := actualValue.(type) {
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
