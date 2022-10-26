package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectFor(a auditor) ast.For {
	n := ast.For{}
	n.Keyword = a.expect(token.For)
	n.Initialiser, n.Condition, n.Advancement = expectForControls(a)
	n.Body = parseStmtBlock(a)
	n.End = a.expect(token.End)
	return n
}

func expectForControls(a auditor) (ast.Stmt, ast.Expr, ast.Stmt) {

	first := acceptInlineStatement(a)
	if first == nil {
		panic(ErrForLoopControls.Track("Expected initialiser or condition"))
	}

	if !a.accept(token.Terminator) {
		expectEndOfStmt(a)
		return nil, forConditionAsExpr(first), nil
	}

	initialiser := first
	condition := expectExpression(a)
	a.expect(token.Terminator)
	advancement := expectStatement(a)

	return initialiser, condition, advancement
}

func forConditionAsExpr(condition ast.Stmt) ast.Expr {
	if ex, ok := condition.(ast.Expr); ok {
		return ex
	}
	panic(ErrForLoopControls.Track("For condition must be an expression"))
}
