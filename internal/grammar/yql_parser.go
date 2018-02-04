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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 24, 106,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 19, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 7, 3, 27, 10, 3, 12, 3, 14, 3, 30, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 56, 10, 4, 12, 4,
	14, 4, 59, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4,
	69, 10, 4, 12, 4, 14, 4, 72, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 7, 4, 82, 10, 4, 12, 4, 14, 4, 85, 11, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 95, 10, 4, 12, 4, 14, 4, 98, 11, 4,
	3, 4, 3, 4, 5, 4, 102, 10, 4, 3, 5, 3, 5, 3, 5, 2, 3, 4, 6, 2, 4, 6, 8,
	2, 3, 4, 2, 18, 19, 21, 23, 2, 117, 2, 10, 3, 2, 2, 2, 4, 18, 3, 2, 2,
	2, 6, 101, 3, 2, 2, 2, 8, 103, 3, 2, 2, 2, 10, 11, 5, 4, 3, 2, 11, 3, 3,
	2, 2, 2, 12, 13, 8, 3, 1, 2, 13, 19, 5, 6, 4, 2, 14, 15, 7, 5, 2, 2, 15,
	16, 5, 4, 3, 2, 16, 17, 7, 6, 2, 2, 17, 19, 3, 2, 2, 2, 18, 12, 3, 2, 2,
	2, 18, 14, 3, 2, 2, 2, 19, 28, 3, 2, 2, 2, 20, 21, 12, 5, 2, 2, 21, 22,
	7, 3, 2, 2, 22, 27, 5, 4, 3, 6, 23, 24, 12, 4, 2, 2, 24, 25, 7, 4, 2, 2,
	25, 27, 5, 4, 3, 5, 26, 20, 3, 2, 2, 2, 26, 23, 3, 2, 2, 2, 27, 30, 3,
	2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 5, 3, 2, 2, 2, 30,
	28, 3, 2, 2, 2, 31, 32, 7, 20, 2, 2, 32, 33, 7, 7, 2, 2, 33, 102, 5, 8,
	5, 2, 34, 35, 7, 20, 2, 2, 35, 36, 7, 8, 2, 2, 36, 102, 5, 8, 5, 2, 37,
	38, 7, 20, 2, 2, 38, 39, 7, 9, 2, 2, 39, 102, 5, 8, 5, 2, 40, 41, 7, 20,
	2, 2, 41, 42, 7, 10, 2, 2, 42, 102, 5, 8, 5, 2, 43, 44, 7, 20, 2, 2, 44,
	45, 7, 11, 2, 2, 45, 102, 5, 8, 5, 2, 46, 47, 7, 20, 2, 2, 47, 48, 7, 12,
	2, 2, 48, 102, 5, 8, 5, 2, 49, 50, 7, 20, 2, 2, 50, 51, 7, 13, 2, 2, 51,
	52, 7, 5, 2, 2, 52, 57, 5, 8, 5, 2, 53, 54, 7, 14, 2, 2, 54, 56, 5, 8,
	5, 2, 55, 53, 3, 2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58,
	3, 2, 2, 2, 58, 60, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 61, 7, 6, 2, 2,
	61, 102, 3, 2, 2, 2, 62, 63, 7, 20, 2, 2, 63, 64, 7, 15, 2, 2, 64, 65,
	7, 5, 2, 2, 65, 70, 5, 8, 5, 2, 66, 67, 7, 14, 2, 2, 67, 69, 5, 8, 5, 2,
	68, 66, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3,
	2, 2, 2, 71, 73, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 74, 7, 6, 2, 2, 74,
	102, 3, 2, 2, 2, 75, 76, 7, 20, 2, 2, 76, 77, 7, 16, 2, 2, 77, 78, 7, 5,
	2, 2, 78, 83, 5, 8, 5, 2, 79, 80, 7, 14, 2, 2, 80, 82, 5, 8, 5, 2, 81,
	79, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2,
	2, 84, 86, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 87, 7, 6, 2, 2, 87, 102,
	3, 2, 2, 2, 88, 89, 7, 20, 2, 2, 89, 90, 7, 17, 2, 2, 90, 91, 7, 5, 2,
	2, 91, 96, 5, 8, 5, 2, 92, 93, 7, 14, 2, 2, 93, 95, 5, 8, 5, 2, 94, 92,
	3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2,
	97, 99, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 100, 7, 6, 2, 2, 100, 102,
	3, 2, 2, 2, 101, 31, 3, 2, 2, 2, 101, 34, 3, 2, 2, 2, 101, 37, 3, 2, 2,
	2, 101, 40, 3, 2, 2, 2, 101, 43, 3, 2, 2, 2, 101, 46, 3, 2, 2, 2, 101,
	49, 3, 2, 2, 2, 101, 62, 3, 2, 2, 2, 101, 75, 3, 2, 2, 2, 101, 88, 3, 2,
	2, 2, 102, 7, 3, 2, 2, 2, 103, 104, 9, 2, 2, 2, 104, 9, 3, 2, 2, 2, 10,
	18, 26, 28, 57, 70, 83, 96, 101,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'and'", "'or'", "'('", "')'", "'='", "'!='", "'>'", "'<'", "'>='",
	"'<='", "'in'", "','", "'!in'", "'\u2229'", "'!\u2229'", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "TRUE",
	"FALSE", "FIELDNAME", "STRING", "INT", "FLOAT", "WS",
}

var ruleNames = []string{
	"query", "expr", "booleanExpr", "value",
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
	YqlParserTRUE      = 16
	YqlParserFALSE     = 17
	YqlParserFIELDNAME = 18
	YqlParserSTRING    = 19
	YqlParserINT       = 20
	YqlParserFLOAT     = 21
	YqlParserWS        = 22
)

// YqlParser rules.
const (
	YqlParserRULE_query       = 0
	YqlParserRULE_expr        = 1
	YqlParserRULE_booleanExpr = 2
	YqlParserRULE_value       = 3
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
		p.SetState(8)
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
	p.SetState(16)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YqlParserFIELDNAME:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(11)
			p.BooleanExpr()
		}

	case YqlParserT__2:
		localctx = NewEmbbedExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(12)
			p.Match(YqlParserT__2)
		}
		{
			p.SetState(13)
			p.expr(0)
		}
		{
			p.SetState(14)
			p.Match(YqlParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(24)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YqlParserRULE_expr)
				p.SetState(18)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(19)
					p.Match(YqlParserT__0)
				}
				{
					p.SetState(20)
					p.expr(4)
				}

			case 2:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YqlParserRULE_expr)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(22)
					p.Match(YqlParserT__1)
				}
				{
					p.SetState(23)
					p.expr(3)
				}

			}

		}
		p.SetState(28)
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

func (s *BooleanExprContext) FIELDNAME() antlr.TerminalNode {
	return s.GetToken(YqlParserFIELDNAME, 0)
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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(YqlParserFIELDNAME)
		}
		{
			p.SetState(30)

			var _m = p.Match(YqlParserT__4)

			localctx.(*BooleanExprContext).op = _m
		}
		{
			p.SetState(31)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Match(YqlParserFIELDNAME)
		}
		{
			p.SetState(33)

			var _m = p.Match(YqlParserT__5)

			localctx.(*BooleanExprContext).op = _m
		}
		{
			p.SetState(34)
			p.Value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.Match(YqlParserFIELDNAME)
		}
		{
			p.SetState(36)

			var _m = p.Match(YqlParserT__6)

			localctx.(*BooleanExprContext).op = _m
		}
		{
			p.SetState(37)
			p.Value()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(38)
			p.Match(YqlParserFIELDNAME)
		}
		{
			p.SetState(39)

			var _m = p.Match(YqlParserT__7)

			localctx.(*BooleanExprContext).op = _m
		}
		{
			p.SetState(40)
			p.Value()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(41)
			p.Match(YqlParserFIELDNAME)
		}
		{
			p.SetState(42)

			var _m = p.Match(YqlParserT__8)

			localctx.(*BooleanExprContext).op = _m
		}
		{
			p.SetState(43)
			p.Value()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(44)
			p.Match(YqlParserFIELDNAME)
		}
		{
			p.SetState(45)

			var _m = p.Match(YqlParserT__9)

			localctx.(*BooleanExprContext).op = _m
		}
		{
			p.SetState(46)
			p.Value()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(47)
			p.Match(YqlParserFIELDNAME)
		}
		{
			p.SetState(48)

			var _m = p.Match(YqlParserT__10)

			localctx.(*BooleanExprContext).op = _m
		}
		{
			p.SetState(49)
			p.Match(YqlParserT__2)
		}
		{
			p.SetState(50)
			p.Value()
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == YqlParserT__11 {
			{
				p.SetState(51)
				p.Match(YqlParserT__11)
			}
			{
				p.SetState(52)
				p.Value()
			}

			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(58)
			p.Match(YqlParserT__3)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(60)
			p.Match(YqlParserFIELDNAME)
		}
		{
			p.SetState(61)

			var _m = p.Match(YqlParserT__12)

			localctx.(*BooleanExprContext).op = _m
		}
		{
			p.SetState(62)
			p.Match(YqlParserT__2)
		}
		{
			p.SetState(63)
			p.Value()
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == YqlParserT__11 {
			{
				p.SetState(64)
				p.Match(YqlParserT__11)
			}
			{
				p.SetState(65)
				p.Value()
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(71)
			p.Match(YqlParserT__3)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(73)
			p.Match(YqlParserFIELDNAME)
		}
		{
			p.SetState(74)

			var _m = p.Match(YqlParserT__13)

			localctx.(*BooleanExprContext).op = _m
		}
		{
			p.SetState(75)
			p.Match(YqlParserT__2)
		}
		{
			p.SetState(76)
			p.Value()
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == YqlParserT__11 {
			{
				p.SetState(77)
				p.Match(YqlParserT__11)
			}
			{
				p.SetState(78)
				p.Value()
			}

			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(84)
			p.Match(YqlParserT__3)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(86)
			p.Match(YqlParserFIELDNAME)
		}
		{
			p.SetState(87)

			var _m = p.Match(YqlParserT__14)

			localctx.(*BooleanExprContext).op = _m
		}
		{
			p.SetState(88)
			p.Match(YqlParserT__2)
		}
		{
			p.SetState(89)
			p.Value()
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == YqlParserT__11 {
			{
				p.SetState(90)
				p.Match(YqlParserT__11)
			}
			{
				p.SetState(91)
				p.Value()
			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.Match(YqlParserT__3)
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
	p.EnterRule(localctx, 6, YqlParserRULE_value)
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
	p.SetState(101)
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
