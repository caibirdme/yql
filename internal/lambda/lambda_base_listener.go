// Generated from Lambda.g4 by ANTLR 4.7.

package lambda // Lambda
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLambdaListener is a complete listener for a parse tree produced by LambdaParser.
type BaseLambdaListener struct{}

var _ LambdaListener = &BaseLambdaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLambdaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLambdaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLambdaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLambdaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseLambdaListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseLambdaListener) ExitLambda(ctx *LambdaContext) {}

// EnterSignature is called when production signature is entered.
func (s *BaseLambdaListener) EnterSignature(ctx *SignatureContext) {}

// ExitSignature is called when production signature is exited.
func (s *BaseLambdaListener) ExitSignature(ctx *SignatureContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseLambdaListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseLambdaListener) ExitOperand(ctx *OperandContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *BaseLambdaListener) EnterBasicLit(ctx *BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *BaseLambdaListener) ExitBasicLit(ctx *BasicLitContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *BaseLambdaListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *BaseLambdaListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseLambdaListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseLambdaListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseLambdaListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseLambdaListener) ExitSelector(ctx *SelectorContext) {}

// EnterFirstExpr is called when production firstExpr is entered.
func (s *BaseLambdaListener) EnterFirstExpr(ctx *FirstExprContext) {}

// ExitFirstExpr is called when production firstExpr is exited.
func (s *BaseLambdaListener) ExitFirstExpr(ctx *FirstExprContext) {}

// EnterValueExpr is called when production valueExpr is entered.
func (s *BaseLambdaListener) EnterValueExpr(ctx *ValueExprContext) {}

// ExitValueExpr is called when production valueExpr is exited.
func (s *BaseLambdaListener) ExitValueExpr(ctx *ValueExprContext) {}

// EnterSecondExpr is called when production secondExpr is entered.
func (s *BaseLambdaListener) EnterSecondExpr(ctx *SecondExprContext) {}

// ExitSecondExpr is called when production secondExpr is exited.
func (s *BaseLambdaListener) ExitSecondExpr(ctx *SecondExprContext) {}

// EnterQuoteExpr is called when production quoteExpr is entered.
func (s *BaseLambdaListener) EnterQuoteExpr(ctx *QuoteExprContext) {}

// ExitQuoteExpr is called when production quoteExpr is exited.
func (s *BaseLambdaListener) ExitQuoteExpr(ctx *QuoteExprContext) {}

// EnterThirdExpr is called when production thirdExpr is entered.
func (s *BaseLambdaListener) EnterThirdExpr(ctx *ThirdExprContext) {}

// ExitThirdExpr is called when production thirdExpr is exited.
func (s *BaseLambdaListener) ExitThirdExpr(ctx *ThirdExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseLambdaListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseLambdaListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseLambdaListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseLambdaListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseLambdaListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseLambdaListener) ExitUnaryExpr(ctx *UnaryExprContext) {}
