package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// FOR := For FOR_CONTROLS STMT_BLOCK End
func acceptFor(a auditor) (ast.For, bool) {
	if !a.accept(token.For) {
		return ast.For{}, false
	}

	n := ast.For{
		Keyword: a.Prev(),
	}

	n.Initialiser, n.Condition, n.Advancement = parseForControls(a)
	n.Body = parseStmtBlock(a)
	n.End = a.expect(token.End)

	return n, true
}

// FOR_CONTROLS := EXPRESSION
// FOR_CONTROLS := [STATEMENT [Terminator EXPRESSION [Terminator STATEMENT]]
func parseForControls(a auditor) (
	initialiser ast.Stmt,
	condition ast.Expr,
	advancement ast.Stmt,
) {
	var ok bool

	initialiser, ok = acceptInlineStatement(a)
	if !a.accept(token.Terminator) {
		expectEndOfStmt(a)

		if ok {
			condition = forConditionAsExpr(initialiser)
			initialiser = nil
		}

		return
	}

	condition, _ = acceptExpression(a)
	if !a.accept(token.Terminator) {
		expectEndOfStmt(a)
		return
	}

	advancement, _ = acceptInlineStatement(a)
	expectEndOfStmt(a)

	return
}

func forConditionAsExpr(condition ast.Stmt) ast.Expr {
	if ex, ok := condition.(ast.Expr); ok {
		return ex
	}
	panic(ErrForLoopControls.Track("For condition must be an expression"))
}
