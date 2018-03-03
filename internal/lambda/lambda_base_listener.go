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

// EnterValueExpr is called when production valueExpr is entered.
func (s *BaseLambdaListener) EnterValueExpr(ctx *ValueExprContext) {}

// ExitValueExpr is called when production valueExpr is exited.
func (s *BaseLambdaListener) ExitValueExpr(ctx *ValueExprContext) {}

// EnterQuoteExpr is called when production quoteExpr is entered.
func (s *BaseLambdaListener) EnterQuoteExpr(ctx *QuoteExprContext) {}

// ExitQuoteExpr is called when production quoteExpr is exited.
func (s *BaseLambdaListener) ExitQuoteExpr(ctx *QuoteExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseLambdaListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseLambdaListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterMulExpr is called when production mulExpr is entered.
func (s *BaseLambdaListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production mulExpr is exited.
func (s *BaseLambdaListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterEqExpre is called when production eqExpre is entered.
func (s *BaseLambdaListener) EnterEqExpre(ctx *EqExpreContext) {}

// ExitEqExpre is called when production eqExpre is exited.
func (s *BaseLambdaListener) ExitEqExpre(ctx *EqExpreContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseLambdaListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseLambdaListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterShiftExpr is called when production shiftExpr is entered.
func (s *BaseLambdaListener) EnterShiftExpr(ctx *ShiftExprContext) {}

// ExitShiftExpr is called when production shiftExpr is exited.
func (s *BaseLambdaListener) ExitShiftExpr(ctx *ShiftExprContext) {}

// EnterCompareExpr is called when production compareExpr is entered.
func (s *BaseLambdaListener) EnterCompareExpr(ctx *CompareExprContext) {}

// ExitCompareExpr is called when production compareExpr is exited.
func (s *BaseLambdaListener) ExitCompareExpr(ctx *CompareExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseLambdaListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseLambdaListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseLambdaListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseLambdaListener) ExitUnaryExpr(ctx *UnaryExprContext) {}
