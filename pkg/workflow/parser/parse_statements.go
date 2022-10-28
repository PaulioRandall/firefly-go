package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrMissingStmt = err.Trackable("Missing statement")

	ErrBadStmt = err.Trackable("Failed to parse statement")
)

// STMT_BLOCK := {STATEMENT}
func parseStmtBlock(a auditor) []ast.Stmt {
	var nodes []ast.Stmt

	for a.isNot(token.End) {
		nodes = append(nodes, expectStatement(a))
	}

	return nodes
}

// INLINE_STATEMENT := [ASSIGNMENT | EXPR]
func acceptInlineStatement(a auditor) (ast.Stmt, bool) {
	defer wrapPanic(func(e error) error {
		return ErrBadStmt.Wrap(e, "Bad inline statement")
	})

	var (
		n  ast.Stmt
		ok bool
	)

	// TODO:
	// - for i, v in expr
	// - spell
	// - func
	// - proc

	if n, ok := acceptAssignment(a); ok {
		return n, true
	}

	if n, ok = acceptExpression(a); ok {
		return n, true
	}

	if n, ok := acceptIf(a); ok {
		return n, true
	}

	if n, ok := acceptFor(a); ok {
		return n, true
	}

	if n, ok := acceptWhen(a); ok {
		return n, true
	}

	if n, ok := acceptWatch(a); ok {
		return n, true
	}

	return nil, false
}

// STATEMENT := [STATEMENT]
func acceptStatement(a auditor) (ast.Stmt, bool) {
	defer wrapPanic(func(e error) error {
		return ErrBadStmt.Wrap(e, "Bad statement")
	})

	if n, ok := acceptInlineStatement(a); ok {
		expectEndOfStmt(a)
		return n, true
	}
	return nil, false
}

// STATEMENT := INLINE_STATEMENT Terminator
func expectStatement(a auditor) ast.Stmt {
	defer wrapPanic(func(e error) error {
		return ErrBadStmt.Wrap(e, "Bad statement")
	})

	if n, ok := acceptStatement(a); ok {
		return n
	}

	panic(ErrMissingStmt.Track("Expected statement"))
}

func expectEndOfStmt(a auditor) {
	if !a.accept(token.Terminator) && !a.accept(token.Newline) {
		panic(a.unexpectedToken("Terminator or newline", a.Peek()))
	}
}
