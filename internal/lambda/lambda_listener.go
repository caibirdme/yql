// Generated from Lambda.g4 by ANTLR 4.7.

package lambda // Lambda
import "github.com/antlr/antlr4/runtime/Go/antlr"

// LambdaListener is a complete listener for a parse tree produced by LambdaParser.
type LambdaListener interface {
	antlr.ParseTreeListener

	// EnterLambda is called when entering the lambda production.
	EnterLambda(c *LambdaContext)

	// EnterSignature is called when entering the signature production.
	EnterSignature(c *SignatureContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterBasicLit is called when entering the basicLit production.
	EnterBasicLit(c *BasicLitContext)

	// EnterOperandName is called when entering the operandName production.
	EnterOperandName(c *OperandNameContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterFirstExpr is called when entering the firstExpr production.
	EnterFirstExpr(c *FirstExprContext)

	// EnterValueExpr is called when entering the valueExpr production.
	EnterValueExpr(c *ValueExprContext)

	// EnterSecondExpr is called when entering the secondExpr production.
	EnterSecondExpr(c *SecondExprContext)

	// EnterQuoteExpr is called when entering the quoteExpr production.
	EnterQuoteExpr(c *QuoteExprContext)

	// EnterThirdExpr is called when entering the thirdExpr production.
	EnterThirdExpr(c *ThirdExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// ExitLambda is called when exiting the lambda production.
	ExitLambda(c *LambdaContext)

	// ExitSignature is called when exiting the signature production.
	ExitSignature(c *SignatureContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitBasicLit is called when exiting the basicLit production.
	ExitBasicLit(c *BasicLitContext)

	// ExitOperandName is called when exiting the operandName production.
	ExitOperandName(c *OperandNameContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitFirstExpr is called when exiting the firstExpr production.
	ExitFirstExpr(c *FirstExprContext)

	// ExitValueExpr is called when exiting the valueExpr production.
	ExitValueExpr(c *ValueExprContext)

	// ExitSecondExpr is called when exiting the secondExpr production.
	ExitSecondExpr(c *SecondExprContext)

	// ExitQuoteExpr is called when exiting the quoteExpr production.
	ExitQuoteExpr(c *QuoteExprContext)

	// ExitThirdExpr is called when exiting the thirdExpr production.
	ExitThirdExpr(c *ThirdExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)
}
