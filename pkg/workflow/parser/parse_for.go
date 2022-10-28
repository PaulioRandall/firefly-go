package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrBadForLoop        = err.Trackable("Failed to parse for loop")
	ErrBadForLoopControl = err.Trackable("Failed to parse for loop controls")
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
	n.End = parseEndOfBlock(a)

	return n, true
}

// FOR_CONTROLS := [EXPRESSION]
func parseForControls(a auditor) (
	initialiser ast.Stmt,
	condition ast.Expr,
	advancement ast.Stmt,
) {
	defer wrapPanic(func(e error) error {
		return ErrBadForLoopControl.Wrap(e, "Bad for loop control syntax")
	})

	var ok bool

	initialiser, ok = acceptInlineStatement(a)
	if a.accept(token.Terminator) {
		return parseNumericForControls(a, initialiser)
	}

	if !ok {
		initialiser = nil
		expectEndOfStmt(a)
		return
	}

	if condition, ok = initialiser.(ast.Expr); ok {
		initialiser = nil
		expectEndOfStmt(a)
	}

	return
}

// FOR_CONTROLS := [STATEMENT] Terminator [EXPRESSION] Terminator [STATEMENT]
func parseNumericForControls(a auditor, initialiser ast.Stmt) (
	ast.Stmt,
	ast.Expr,
	ast.Stmt,
) {
	condition, _ := acceptExpression(a)
	parseTerminator(a)
	advancement, _ := acceptStatement(a)

	if advancement == nil {
		expectEndOfStmt(a)
	}

	return initialiser, condition, advancement
}
