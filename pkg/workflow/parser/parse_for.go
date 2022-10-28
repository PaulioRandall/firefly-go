package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrBadForLoop = err.Trackable("Failed to parse for loop")
)

// FOR := For FOR_CONTROLS STMT_BLOCK End
func acceptFor(a auditor) (ast.For, bool) {
	defer wrapPanic(func(e error) error {
		return ErrBadForLoop.Wrap(e, "Bad for loop syntax")
	})

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

// FOR_CONTROLS := [EXPRESSION]
// FOR_CONTROLS := [STATEMENT] Terminator [EXPRESSION] Terminator [STATEMENT]
func parseForControls(a auditor) (
	initialiser ast.Stmt,
	condition ast.Expr,
	advancement ast.Stmt,
) {
	var ok bool

	if initialiser, ok = acceptInlineStatement(a); !ok {
		initialiser = nil
		expectEndOfStmt(a)
		return
	}

	if condition, ok = initialiser.(ast.Expr); ok {
		initialiser = nil
		expectEndOfStmt(a)
		return
	}

	a.expect(token.Terminator)
	condition, _ = acceptExpression(a)
	a.expect(token.Terminator)
	advancement, _ = acceptStatement(a)

	return
}
