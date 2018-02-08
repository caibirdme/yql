// Generated from Yql.g4 by ANTLR 4.7.

package grammar // Yql
import "github.com/antlr/antlr4/runtime/Go/antlr"

// YqlListener is a complete listener for a parse tree produced by YqlParser.
type YqlListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterEmbbedExpr is called when entering the embbedExpr production.
	EnterEmbbedExpr(c *EmbbedExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterBoolExpr is called when entering the boolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterBooleanExpr is called when entering the booleanExpr production.
	EnterBooleanExpr(c *BooleanExprContext)

	// EnterLeftexpr is called when entering the leftexpr production.
	EnterLeftexpr(c *LeftexprContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitEmbbedExpr is called when exiting the embbedExpr production.
	ExitEmbbedExpr(c *EmbbedExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitBoolExpr is called when exiting the boolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitBooleanExpr is called when exiting the booleanExpr production.
	ExitBooleanExpr(c *BooleanExprContext)

	// ExitLeftexpr is called when exiting the leftexpr production.
	ExitLeftexpr(c *LeftexprContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
