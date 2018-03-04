// Generated from Lambda.g4 by ANTLR 4.7.

package lambda // Lambda
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 90, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 5, 3, 34, 10, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 5, 4, 40, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	7, 6, 47, 10, 6, 12, 6, 14, 6, 50, 11, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 63, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	7, 9, 80, 10, 9, 12, 9, 14, 9, 83, 11, 9, 3, 10, 3, 10, 3, 10, 5, 10, 88,
	10, 10, 3, 10, 2, 3, 16, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 7, 4, 2,
	31, 35, 38, 38, 3, 2, 8, 14, 3, 2, 15, 18, 3, 2, 19, 24, 7, 2, 8, 8, 13,
	13, 15, 16, 18, 18, 27, 28, 2, 91, 2, 20, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2,
	6, 39, 3, 2, 2, 2, 8, 41, 3, 2, 2, 2, 10, 43, 3, 2, 2, 2, 12, 51, 3, 2,
	2, 2, 14, 53, 3, 2, 2, 2, 16, 62, 3, 2, 2, 2, 18, 87, 3, 2, 2, 2, 20, 21,
	5, 4, 3, 2, 21, 22, 7, 3, 2, 2, 22, 23, 5, 16, 9, 2, 23, 3, 3, 2, 2, 2,
	24, 33, 7, 4, 2, 2, 25, 30, 7, 29, 2, 2, 26, 27, 7, 5, 2, 2, 27, 29, 7,
	29, 2, 2, 28, 26, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30,
	31, 3, 2, 2, 2, 31, 34, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33, 25, 3, 2, 2,
	2, 33, 34, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 36, 7, 6, 2, 2, 36, 5, 3,
	2, 2, 2, 37, 40, 5, 8, 5, 2, 38, 40, 5, 10, 6, 2, 39, 37, 3, 2, 2, 2, 39,
	38, 3, 2, 2, 2, 40, 7, 3, 2, 2, 2, 41, 42, 9, 2, 2, 2, 42, 9, 3, 2, 2,
	2, 43, 48, 7, 29, 2, 2, 44, 45, 7, 7, 2, 2, 45, 47, 7, 29, 2, 2, 46, 44,
	3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2,
	49, 11, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 52, 5, 6, 4, 2, 52, 13, 3,
	2, 2, 2, 53, 54, 7, 7, 2, 2, 54, 55, 7, 29, 2, 2, 55, 15, 3, 2, 2, 2, 56,
	57, 8, 9, 1, 2, 57, 58, 7, 4, 2, 2, 58, 59, 5, 16, 9, 2, 59, 60, 7, 6,
	2, 2, 60, 63, 3, 2, 2, 2, 61, 63, 5, 18, 10, 2, 62, 56, 3, 2, 2, 2, 62,
	61, 3, 2, 2, 2, 63, 81, 3, 2, 2, 2, 64, 65, 12, 8, 2, 2, 65, 66, 9, 3,
	2, 2, 66, 80, 5, 16, 9, 9, 67, 68, 12, 7, 2, 2, 68, 69, 9, 4, 2, 2, 69,
	80, 5, 16, 9, 8, 70, 71, 12, 6, 2, 2, 71, 72, 9, 5, 2, 2, 72, 80, 5, 16,
	9, 7, 73, 74, 12, 5, 2, 2, 74, 75, 7, 25, 2, 2, 75, 80, 5, 16, 9, 6, 76,
	77, 12, 4, 2, 2, 77, 78, 7, 26, 2, 2, 78, 80, 5, 16, 9, 5, 79, 64, 3, 2,
	2, 2, 79, 67, 3, 2, 2, 2, 79, 70, 3, 2, 2, 2, 79, 73, 3, 2, 2, 2, 79, 76,
	3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2,
	82, 17, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 88, 5, 12, 7, 2, 85, 86, 9,
	6, 2, 2, 86, 88, 5, 18, 10, 2, 87, 84, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2,
	88, 19, 3, 2, 2, 2, 10, 30, 33, 39, 48, 62, 79, 81, 87,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'=>'", "'('", "','", "')'", "'.'", "'*'", "'/'", "'%'", "'<<'", "'>>'",
	"'&'", "'&^'", "'+'", "'-'", "'|'", "'^'", "'=='", "'!='", "'<'", "'<='",
	"'>'", "'>='", "'&&'", "'||'", "'!'", "'<-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "IDENTIFIER", "BINARY_OP", "INT_LIT",
	"BOOL_LIT", "FLOAT_LIT", "IMAGINARY_LIT", "RUNE_LIT", "LITTLE_U_VALUE",
	"BIG_U_VALUE", "STRING_LIT", "WS",
}

var ruleNames = []string{
	"lambda", "signature", "operand", "basicLit", "operandName", "primaryExpr",
	"selector", "expression", "unaryExpr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type LambdaParser struct {
	*antlr.BaseParser
}

func NewLambdaParser(input antlr.TokenStream) *LambdaParser {
	this := new(LambdaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Lambda.g4"

	return this
}

// LambdaParser tokens.
const (
	LambdaParserEOF            = antlr.TokenEOF
	LambdaParserT__0           = 1
	LambdaParserT__1           = 2
	LambdaParserT__2           = 3
	LambdaParserT__3           = 4
	LambdaParserT__4           = 5
	LambdaParserT__5           = 6
	LambdaParserT__6           = 7
	LambdaParserT__7           = 8
	LambdaParserT__8           = 9
	LambdaParserT__9           = 10
	LambdaParserT__10          = 11
	LambdaParserT__11          = 12
	LambdaParserT__12          = 13
	LambdaParserT__13          = 14
	LambdaParserT__14          = 15
	LambdaParserT__15          = 16
	LambdaParserT__16          = 17
	LambdaParserT__17          = 18
	LambdaParserT__18          = 19
	LambdaParserT__19          = 20
	LambdaParserT__20          = 21
	LambdaParserT__21          = 22
	LambdaParserT__22          = 23
	LambdaParserT__23          = 24
	LambdaParserT__24          = 25
	LambdaParserT__25          = 26
	LambdaParserIDENTIFIER     = 27
	LambdaParserBINARY_OP      = 28
	LambdaParserINT_LIT        = 29
	LambdaParserBOOL_LIT       = 30
	LambdaParserFLOAT_LIT      = 31
	LambdaParserIMAGINARY_LIT  = 32
	LambdaParserRUNE_LIT       = 33
	LambdaParserLITTLE_U_VALUE = 34
	LambdaParserBIG_U_VALUE    = 35
	LambdaParserSTRING_LIT     = 36
	LambdaParserWS             = 37
)

// LambdaParser rules.
const (
	LambdaParserRULE_lambda      = 0
	LambdaParserRULE_signature   = 1
	LambdaParserRULE_operand     = 2
	LambdaParserRULE_basicLit    = 3
	LambdaParserRULE_operandName = 4
	LambdaParserRULE_primaryExpr = 5
	LambdaParserRULE_selector    = 6
	LambdaParserRULE_expression  = 7
	LambdaParserRULE_unaryExpr   = 8
)

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) Signature() ISignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignatureContext)
}

func (s *LambdaContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (p *LambdaParser) Lambda() (localctx ILambdaContext) {
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LambdaParserRULE_lambda)

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
		p.SetState(18)
		p.Signature()
	}
	{
		p.SetState(19)
		p.Match(LambdaParserT__0)
	}
	{
		p.SetState(20)
		p.expression(0)
	}

	return localctx
}

// ISignatureContext is an interface to support dynamic dispatch.
type ISignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignatureContext differentiates from other interfaces.
	IsSignatureContext()
}

type SignatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignatureContext() *SignatureContext {
	var p = new(SignatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_signature
	return p
}

func (*SignatureContext) IsSignatureContext() {}

func NewSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignatureContext {
	var p = new(SignatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_signature

	return p
}

func (s *SignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *SignatureContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(LambdaParserIDENTIFIER)
}

func (s *SignatureContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(LambdaParserIDENTIFIER, i)
}

func (s *SignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterSignature(s)
	}
}

func (s *SignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitSignature(s)
	}
}

func (p *LambdaParser) Signature() (localctx ISignatureContext) {
	localctx = NewSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LambdaParserRULE_signature)
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
		p.SetState(22)
		p.Match(LambdaParserT__1)
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LambdaParserIDENTIFIER {
		{
			p.SetState(23)
			p.Match(LambdaParserIDENTIFIER)
		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LambdaParserT__2 {
			{
				p.SetState(24)
				p.Match(LambdaParserT__2)
			}
			{
				p.SetState(25)
				p.Match(LambdaParserIDENTIFIER)
			}

			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(33)
		p.Match(LambdaParserT__3)
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) BasicLit() IBasicLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (p *LambdaParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LambdaParserRULE_operand)

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

	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LambdaParserINT_LIT, LambdaParserBOOL_LIT, LambdaParserFLOAT_LIT, LambdaParserIMAGINARY_LIT, LambdaParserRUNE_LIT, LambdaParserSTRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.BasicLit()
		}

	case LambdaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.OperandName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBasicLitContext is an interface to support dynamic dispatch.
type IBasicLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicLitContext differentiates from other interfaces.
	IsBasicLitContext()
}

type BasicLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicLitContext() *BasicLitContext {
	var p = new(BasicLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(LambdaParserINT_LIT, 0)
}

func (s *BasicLitContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(LambdaParserBOOL_LIT, 0)
}

func (s *BasicLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(LambdaParserFLOAT_LIT, 0)
}

func (s *BasicLitContext) IMAGINARY_LIT() antlr.TerminalNode {
	return s.GetToken(LambdaParserIMAGINARY_LIT, 0)
}

func (s *BasicLitContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(LambdaParserRUNE_LIT, 0)
}

func (s *BasicLitContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(LambdaParserSTRING_LIT, 0)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterBasicLit(s)
	}
}

func (s *BasicLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitBasicLit(s)
	}
}

func (p *LambdaParser) BasicLit() (localctx IBasicLitContext) {
	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LambdaParserRULE_basicLit)
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
	p.SetState(39)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(LambdaParserINT_LIT-29))|(1<<(LambdaParserBOOL_LIT-29))|(1<<(LambdaParserFLOAT_LIT-29))|(1<<(LambdaParserIMAGINARY_LIT-29))|(1<<(LambdaParserRUNE_LIT-29))|(1<<(LambdaParserSTRING_LIT-29)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_operandName
	return p
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(LambdaParserIDENTIFIER)
}

func (s *OperandNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(LambdaParserIDENTIFIER, i)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterOperandName(s)
	}
}

func (s *OperandNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitOperandName(s)
	}
}

func (p *LambdaParser) OperandName() (localctx IOperandNameContext) {
	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LambdaParserRULE_operandName)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(LambdaParserIDENTIFIER)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(42)
				p.Match(LambdaParserT__4)
			}
			{
				p.SetState(43)
				p.Match(LambdaParserIDENTIFIER)
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (p *LambdaParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LambdaParserRULE_primaryExpr)

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
		p.SetState(49)
		p.Operand()
	}

	return localctx
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_selector
	return p
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LambdaParserIDENTIFIER, 0)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (p *LambdaParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LambdaParserRULE_selector)

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
		p.SetState(51)
		p.Match(LambdaParserT__4)
	}
	{
		p.SetState(52)
		p.Match(LambdaParserIDENTIFIER)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FirstExprContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewFirstExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FirstExprContext {
	var p = new(FirstExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FirstExprContext) GetOp() antlr.Token { return s.op }

func (s *FirstExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *FirstExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FirstExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FirstExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FirstExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterFirstExpr(s)
	}
}

func (s *FirstExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitFirstExpr(s)
	}
}

type ValueExprContext struct {
	*ExpressionContext
}

func NewValueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueExprContext {
	var p = new(ValueExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) UnaryExpr() IUnaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *ValueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterValueExpr(s)
	}
}

func (s *ValueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitValueExpr(s)
	}
}

type SecondExprContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewSecondExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SecondExprContext {
	var p = new(SecondExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SecondExprContext) GetOp() antlr.Token { return s.op }

func (s *SecondExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *SecondExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecondExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SecondExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SecondExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterSecondExpr(s)
	}
}

func (s *SecondExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitSecondExpr(s)
	}
}

type QuoteExprContext struct {
	*ExpressionContext
}

func NewQuoteExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuoteExprContext {
	var p = new(QuoteExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *QuoteExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuoteExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *QuoteExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterQuoteExpr(s)
	}
}

func (s *QuoteExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitQuoteExpr(s)
	}
}

type ThirdExprContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewThirdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThirdExprContext {
	var p = new(ThirdExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ThirdExprContext) GetOp() antlr.Token { return s.op }

func (s *ThirdExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ThirdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThirdExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ThirdExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ThirdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterThirdExpr(s)
	}
}

func (s *ThirdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitThirdExpr(s)
	}
}

type OrExprContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrExprContext) GetOp() antlr.Token { return s.op }

func (s *OrExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

type AndExprContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndExprContext) GetOp() antlr.Token { return s.op }

func (s *AndExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (p *LambdaParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *LambdaParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, LambdaParserRULE_expression, _p)
	var _la int

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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LambdaParserT__1:
		localctx = NewQuoteExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(55)
			p.Match(LambdaParserT__1)
		}
		{
			p.SetState(56)
			p.expression(0)
		}
		{
			p.SetState(57)
			p.Match(LambdaParserT__3)
		}

	case LambdaParserT__5, LambdaParserT__10, LambdaParserT__12, LambdaParserT__13, LambdaParserT__15, LambdaParserT__24, LambdaParserT__25, LambdaParserIDENTIFIER, LambdaParserINT_LIT, LambdaParserBOOL_LIT, LambdaParserFLOAT_LIT, LambdaParserIMAGINARY_LIT, LambdaParserRUNE_LIT, LambdaParserSTRING_LIT:
		localctx = NewValueExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)
			p.UnaryExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFirstExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LambdaParserRULE_expression)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				p.SetState(63)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*FirstExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LambdaParserT__5)|(1<<LambdaParserT__6)|(1<<LambdaParserT__7)|(1<<LambdaParserT__8)|(1<<LambdaParserT__9)|(1<<LambdaParserT__10)|(1<<LambdaParserT__11))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*FirstExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(64)
					p.expression(7)
				}

			case 2:
				localctx = NewSecondExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LambdaParserRULE_expression)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(66)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*SecondExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LambdaParserT__12)|(1<<LambdaParserT__13)|(1<<LambdaParserT__14)|(1<<LambdaParserT__15))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*SecondExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(67)
					p.expression(6)
				}

			case 3:
				localctx = NewThirdExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LambdaParserRULE_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(69)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ThirdExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LambdaParserT__16)|(1<<LambdaParserT__17)|(1<<LambdaParserT__18)|(1<<LambdaParserT__19)|(1<<LambdaParserT__20)|(1<<LambdaParserT__21))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ThirdExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(70)
					p.expression(5)
				}

			case 4:
				localctx = NewAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LambdaParserRULE_expression)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(72)

					var _m = p.Match(LambdaParserT__22)

					localctx.(*AndExprContext).op = _m
				}
				{
					p.SetState(73)
					p.expression(4)
				}

			case 5:
				localctx = NewOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LambdaParserRULE_expression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(75)

					var _m = p.Match(LambdaParserT__23)

					localctx.(*OrExprContext).op = _m
				}
				{
					p.SetState(76)
					p.expression(3)
				}

			}

		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_unaryExpr
	return p
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *UnaryExprContext) UnaryExpr() IUnaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (p *LambdaParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LambdaParserRULE_unaryExpr)
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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LambdaParserIDENTIFIER, LambdaParserINT_LIT, LambdaParserBOOL_LIT, LambdaParserFLOAT_LIT, LambdaParserIMAGINARY_LIT, LambdaParserRUNE_LIT, LambdaParserSTRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.PrimaryExpr()
		}

	case LambdaParserT__5, LambdaParserT__10, LambdaParserT__12, LambdaParserT__13, LambdaParserT__15, LambdaParserT__24, LambdaParserT__25:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(83)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LambdaParserT__5)|(1<<LambdaParserT__10)|(1<<LambdaParserT__12)|(1<<LambdaParserT__13)|(1<<LambdaParserT__15)|(1<<LambdaParserT__24)|(1<<LambdaParserT__25))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(84)
			p.UnaryExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *LambdaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LambdaParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
