package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrBadLoop           = err.Trackable("Failed to parse loop")
	ErrBadForLoop        = err.Trackable("Failed to parse for loop")
	ErrBadForEachLoop    = err.Trackable("Failed to parse for each loop")
	ErrBadForLoopControl = err.Trackable("Failed to parse for loop controls")
)

// LOOP := FOR | FOR_EACH
func acceptLoop(a auditor) (ast.Stmt, bool) {
	defer wrapPanic(func(e error) error {
		return ErrBadLoop.Wrap(e, "Bad loop")
	})

	if n, ok := acceptFor(a); ok {
		return n, true
	}

	if n, ok := acceptForEach(a); ok {
		return n, true
	}

	return nil, false
}

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
// FOR_CONTROLS := [STATEMENT] Terminator [EXPRESSION] Terminator [STATEMENT]
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
		return parseIteratingForControls(a, initialiser)
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

func parseIteratingForControls(a auditor, initialiser ast.Stmt) (
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

// FOR_EACH := For FOR_EACH_CONTROLS STMT_BLOCK End
func acceptForEach(a auditor) (ast.ForEach, bool) {
	// TODO: Write tests first
	return ast.ForEach{}, false
}

// FOR_EACH_CONTROLS := Variable Comma Variable In EXPRESSION
