package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var MissingStmt = err.Trackable("Missing statement")

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

	if n, ok := acceptFor(a); ok {
		return n, true
	}

	switch {
	case a.is(token.If):
		return expectIf(a), true

	case a.is(token.When):
		return expectWhen(a), true

	case a.is(token.Watch):
		return expectWatch(a), true

	default:
		return nil, false
	}
}

// STATEMENT := [STATEMENT]
func acceptStatement(a auditor) (ast.Stmt, bool) {
	if n, ok := acceptInlineStatement(a); ok {
		expectEndOfStmt(a)
		return n, true
	}
	return nil, false
}

// STATEMENT := INLINE_STATEMENT Terminator
func expectStatement(a auditor) ast.Stmt {
	if n, ok := acceptStatement(a); ok {
		return n
	}
	panic(MissingStmt.Track("Expected statement"))
}

func expectEndOfStmt(a auditor) {
	if !a.accept(token.Terminator) && !a.accept(token.Newline) {
		panic(a.unexpectedToken("Terminator or newline", a.Peek()))
	}
}
