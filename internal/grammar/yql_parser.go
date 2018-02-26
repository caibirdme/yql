// Generated from Yql.g4 by ANTLR 4.7.

package grammar // Yql
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 65, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 21, 10, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 44, 10, 4, 12, 4, 14,
	4, 47, 11, 4, 3, 4, 3, 4, 5, 4, 51, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5,
	57, 10, 5, 13, 5, 14, 5, 58, 5, 5, 61, 10, 5, 3, 6, 3, 6, 3, 6, 2, 3, 4,
	7, 2, 4, 6, 8, 10, 2, 5, 3, 2, 7, 12, 3, 2, 13, 16, 4, 2, 20, 21, 24, 26,
	2, 66, 2, 12, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 50, 3, 2, 2, 2, 8, 52,
	3, 2, 2, 2, 10, 62, 3, 2, 2, 2, 12, 13, 5, 4, 3, 2, 13, 3, 3, 2, 2, 2,
	14, 15, 8, 3, 1, 2, 15, 21, 5, 6, 4, 2, 16, 17, 7, 5, 2, 2, 17, 18, 5,
	4, 3, 2, 18, 19, 7, 6, 2, 2, 19, 21, 3, 2, 2, 2, 20, 14, 3, 2, 2, 2, 20,
	16, 3, 2, 2, 2, 21, 30, 3, 2, 2, 2, 22, 23, 12, 5, 2, 2, 23, 24, 7, 3,
	2, 2, 24, 29, 5, 4, 3, 6, 25, 26, 12, 4, 2, 2, 26, 27, 7, 4, 2, 2, 27,
	29, 5, 4, 3, 5, 28, 22, 3, 2, 2, 2, 28, 25, 3, 2, 2, 2, 29, 32, 3, 2, 2,
	2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 30, 3,
	2, 2, 2, 33, 34, 5, 8, 5, 2, 34, 35, 9, 2, 2, 2, 35, 36, 5, 10, 6, 2, 36,
	51, 3, 2, 2, 2, 37, 38, 5, 8, 5, 2, 38, 39, 9, 3, 2, 2, 39, 40, 7, 5, 2,
	2, 40, 45, 5, 10, 6, 2, 41, 42, 7, 17, 2, 2, 42, 44, 5, 10, 6, 2, 43, 41,
	3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2,
	46, 48, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 49, 7, 6, 2, 2, 49, 51, 3,
	2, 2, 2, 50, 33, 3, 2, 2, 2, 50, 37, 3, 2, 2, 2, 51, 7, 3, 2, 2, 2, 52,
	60, 7, 23, 2, 2, 53, 54, 7, 18, 2, 2, 54, 55, 7, 22, 2, 2, 55, 57, 7, 19,
	2, 2, 56, 53, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59,
	3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 56, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2,
	61, 9, 3, 2, 2, 2, 62, 63, 9, 4, 2, 2, 63, 11, 3, 2, 2, 2, 9, 20, 28, 30,
	45, 50, 58, 60,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'and'", "'or'", "'('", "')'", "'='", "'!='", "'>'", "'<'", "'>='",
	"'<='", "'in'", "'!in'", "'\u2229'", "'!\u2229'", "','", "'.'", "'()'",
	"'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"TRUE", "FALSE", "FUNC", "FIELDNAME", "STRING", "INT", "FLOAT", "WS",
}

var ruleNames = []string{
	"query", "expr", "booleanExpr", "leftexpr", "value",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type YqlParser struct {
	*antlr.BaseParser
}

func NewYqlParser(input antlr.TokenStream) *YqlParser {
	this := new(YqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Yql.g4"

	return this
}

// YqlParser tokens.
const (
	YqlParserEOF       = antlr.TokenEOF
	YqlParserT__0      = 1
	YqlParserT__1      = 2
	YqlParserT__2      = 3
	YqlParserT__3      = 4
	YqlParserT__4      = 5
	YqlParserT__5      = 6
	YqlParserT__6      = 7
	YqlParserT__7      = 8
	YqlParserT__8      = 9
	YqlParserT__9      = 10
	YqlParserT__10     = 11
	YqlParserT__11     = 12
	YqlParserT__12     = 13
	YqlParserT__13     = 14
	YqlParserT__14     = 15
	YqlParserT__15     = 16
	YqlParserT__16     = 17
	YqlParserTRUE      = 18
	YqlParserFALSE     = 19
	YqlParserFUNC      = 20
	YqlParserFIELDNAME = 21
	YqlParserSTRING    = 22
	YqlParserINT       = 23
	YqlParserFLOAT     = 24
	YqlParserWS        = 25
)

// YqlParser rules.
const (
	YqlParserRULE_query       = 0
	YqlParserRULE_expr        = 1
	YqlParserRULE_booleanExpr = 2
	YqlParserRULE_leftexpr    = 3
	YqlParserRULE_value       = 4
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YqlParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YqlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *YqlParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YqlParserRULE_query)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.expr(0)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YqlParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YqlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EmbbedExprContext struct {
	*ExprContext
}

func NewEmbbedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmbbedExprContext {
	var p = new(EmbbedExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EmbbedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmbbedExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EmbbedExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.EnterEmbbedExpr(s)
	}
}

func (s *EmbbedExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.ExitEmbbedExpr(s)
	}
}

type OrExprContext struct {
	*ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

type BoolExprContext struct {
	*ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) BooleanExpr() IBooleanExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanExprContext)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

type AndExprContext struct {
	*ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (p *YqlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *YqlParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, YqlParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YqlParserFIELDNAME:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(13)
			p.BooleanExpr()
		}

	case YqlParserT__2:
		localctx = NewEmbbedExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(14)
			p.Match(YqlParserT__2)
		}
		{
			p.SetState(15)
			p.expr(0)
		}
		{
			p.SetState(16)
			p.Match(YqlParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YqlParserRULE_expr)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(21)
					p.Match(YqlParserT__0)
				}
				{
					p.SetState(22)
					p.expr(4)
				}

			case 2:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YqlParserRULE_expr)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(24)
					p.Match(YqlParserT__1)
				}
				{
					p.SetState(25)
					p.expr(3)
				}

			}

		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IBooleanExprContext is an interface to support dynamic dispatch.
type IBooleanExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsBooleanExprContext differentiates from other interfaces.
	IsBooleanExprContext()
}

type BooleanExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyBooleanExprContext() *BooleanExprContext {
	var p = new(BooleanExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YqlParserRULE_booleanExpr
	return p
}

func (*BooleanExprContext) IsBooleanExprContext() {}

func NewBooleanExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanExprContext {
	var p = new(BooleanExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YqlParserRULE_booleanExpr

	return p
}

func (s *BooleanExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanExprContext) GetOp() antlr.Token { return s.op }

func (s *BooleanExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BooleanExprContext) Leftexpr() ILeftexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeftexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *BooleanExprContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *BooleanExprContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *BooleanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.EnterBooleanExpr(s)
	}
}

func (s *BooleanExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.ExitBooleanExpr(s)
	}
}

func (p *YqlParser) BooleanExpr() (localctx IBooleanExprContext) {
	localctx = NewBooleanExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, YqlParserRULE_booleanExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Leftexpr()
		}
		p.SetState(32)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*BooleanExprContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<YqlParserT__4)|(1<<YqlParserT__5)|(1<<YqlParserT__6)|(1<<YqlParserT__7)|(1<<YqlParserT__8)|(1<<YqlParserT__9))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*BooleanExprContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(33)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Leftexpr()
		}
		p.SetState(36)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*BooleanExprContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<YqlParserT__10)|(1<<YqlParserT__11)|(1<<YqlParserT__12)|(1<<YqlParserT__13))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*BooleanExprContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(37)
			p.Match(YqlParserT__2)
		}
		{
			p.SetState(38)
			p.Value()
		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == YqlParserT__14 {
			{
				p.SetState(39)
				p.Match(YqlParserT__14)
			}
			{
				p.SetState(40)
				p.Value()
			}

			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(46)
			p.Match(YqlParserT__3)
		}

	}

	return localctx
}

// ILeftexprContext is an interface to support dynamic dispatch.
type ILeftexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeftexprContext differentiates from other interfaces.
	IsLeftexprContext()
}

type LeftexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeftexprContext() *LeftexprContext {
	var p = new(LeftexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YqlParserRULE_leftexpr
	return p
}

func (*LeftexprContext) IsLeftexprContext() {}

func NewLeftexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftexprContext {
	var p = new(LeftexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YqlParserRULE_leftexpr

	return p
}

func (s *LeftexprContext) GetParser() antlr.Parser { return s.parser }

func (s *LeftexprContext) FIELDNAME() antlr.TerminalNode {
	return s.GetToken(YqlParserFIELDNAME, 0)
}

func (s *LeftexprContext) AllFUNC() []antlr.TerminalNode {
	return s.GetTokens(YqlParserFUNC)
}

func (s *LeftexprContext) FUNC(i int) antlr.TerminalNode {
	return s.GetToken(YqlParserFUNC, i)
}

func (s *LeftexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.EnterLeftexpr(s)
	}
}

func (s *LeftexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.ExitLeftexpr(s)
	}
}

func (p *YqlParser) Leftexpr() (localctx ILeftexprContext) {
	localctx = NewLeftexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, YqlParserRULE_leftexpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(YqlParserFIELDNAME)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YqlParserT__15 {
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == YqlParserT__15 {
			{
				p.SetState(51)
				p.Match(YqlParserT__15)
			}
			{
				p.SetState(52)
				p.Match(YqlParserFUNC)
			}
			{
				p.SetState(53)
				p.Match(YqlParserT__16)
			}

			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YqlParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YqlParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(YqlParserSTRING, 0)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(YqlParserINT, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(YqlParserFLOAT, 0)
}

func (s *ValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(YqlParserTRUE, 0)
}

func (s *ValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(YqlParserFALSE, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YqlListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *YqlParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, YqlParserRULE_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(60)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<YqlParserTRUE)|(1<<YqlParserFALSE)|(1<<YqlParserSTRING)|(1<<YqlParserINT)|(1<<YqlParserFLOAT))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

func (p *YqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *YqlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
