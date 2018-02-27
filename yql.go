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
}

type yqlListener struct {
	*grammar.BaseYqlListener
	stack       boolStack
	data        map[string]interface{}
	fieldName   string
	funcs       []string
	notFoundErr error
}

func (l *yqlListener) ExitBooleanExpr(ctx *grammar.BooleanExprContext) {
	if l.notFoundErr != nil {
		return
	}
	operator := ctx.GetOp().GetText()
	fieldName := l.fieldName
	actualValue, ok := l.data[fieldName]
	if !ok {
		l.notFoundErr = fmt.Errorf("%s not provided", fieldName)
		return
	}
	var err error
	actualValue, err = funcRuner(actualValue, l.funcs)
	if err != nil {
		l.notFoundErr = err
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

func (l *yqlListener) ExitLeftexpr(ctx *grammar.LeftexprContext) {
	l.fieldName = ctx.FIELDNAME().GetText()
	funcs := ctx.AllFUNC()
	l.funcs = l.funcs[:0]
	if 0 == len(funcs) {
		return
	}
	for _, f := range funcs {
		l.funcs = append(l.funcs, f.GetText())
	}
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

type bailLexer struct {
	*grammar.YqlLexer
}

// Override default Recover and do nothing
// lexer can quit when any error occur
func (l *bailLexer) Recover(r antlr.RecognitionException) {
	panic(r)
}

var (
	bailErrStrategy = antlr.NewBailErrorStrategy()
)

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
	lexer := &bailLexer{grammar.NewYqlLexer(inputStream)}
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewYqlParser(stream)
	parser.BuildParseTrees = true
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(bailErrStrategy)
	tree := parser.Query()
	s, release := stack.NewStack()
	defer release()
	l := &yqlListener{stack: s, funcs: make([]string, 0, 1)}
	l.data = data
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	if l.notFoundErr != nil {
		return false, l.notFoundErr
	}
	return l.stack.Pop(), nil
}

// Ruler caches an AST
type cachedAST struct {
	query grammar.IQueryContext
}

// Ruler represents an AST structure and could do the match according to input data
type Ruler interface {
	Match(map[string]interface{}) (bool, error)
}

// Match do the comparison according to the input data
func (ast cachedAST) Match(data map[string]interface{}) (bool, error) {
	if 0 == len(data) {
		return false, nil
	}
	s, release := stack.NewStack()
	defer release()
	l := &yqlListener{stack: s, funcs: make([]string, 0, 1)}
	l.data = data
	antlr.ParseTreeWalkerDefault.Walk(l, ast.query)
	if l.notFoundErr != nil {
		return false, l.notFoundErr
	}
	return l.stack.Pop(), nil
}

// Rule analyze a rule and store the AST
// it returns error when receives illegal yql expression
// It's more faster than using Match directly
func Rule(rawYQL string) (ruler Ruler, err error) {
	ast := cachedAST{}
	defer func() {
		if r := recover(); nil != r {
			err = fmt.Errorf("%v", r)
			ruler = ast
		}
	}()
	inputStream := antlr.NewInputStream(rawYQL)
	lexer := &bailLexer{grammar.NewYqlLexer(inputStream)}
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewYqlParser(stream)
	parser.BuildParseTrees = true
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(bailErrStrategy)
	ast.query = parser.Query()
	return ast, nil
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
	case bool:
		expect, err := strconv.ParseBool(e)
		if nil != err {
			return false
		}
		return cmpBool(actual, expect, op)
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
