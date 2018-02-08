// Generated from Yql.g4 by ANTLR 4.7.

package grammar // Yql
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseYqlListener is a complete listener for a parse tree produced by YqlParser.
type BaseYqlListener struct{}

var _ YqlListener = &BaseYqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseYqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseYqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseYqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseYqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseYqlListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseYqlListener) ExitQuery(ctx *QueryContext) {}

// EnterEmbbedExpr is called when production embbedExpr is entered.
func (s *BaseYqlListener) EnterEmbbedExpr(ctx *EmbbedExprContext) {}

// ExitEmbbedExpr is called when production embbedExpr is exited.
func (s *BaseYqlListener) ExitEmbbedExpr(ctx *EmbbedExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseYqlListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseYqlListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterBoolExpr is called when production boolExpr is entered.
func (s *BaseYqlListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production boolExpr is exited.
func (s *BaseYqlListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseYqlListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseYqlListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterBooleanExpr is called when production booleanExpr is entered.
func (s *BaseYqlListener) EnterBooleanExpr(ctx *BooleanExprContext) {}

// ExitBooleanExpr is called when production booleanExpr is exited.
func (s *BaseYqlListener) ExitBooleanExpr(ctx *BooleanExprContext) {}

// EnterLeftexpr is called when production leftexpr is entered.
func (s *BaseYqlListener) EnterLeftexpr(ctx *LeftexprContext) {}

// ExitLeftexpr is called when production leftexpr is exited.
func (s *BaseYqlListener) ExitLeftexpr(ctx *LeftexprContext) {}

// EnterValue is called when production value is entered.
func (s *BaseYqlListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseYqlListener) ExitValue(ctx *ValueContext) {}
