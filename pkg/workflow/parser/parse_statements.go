package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func parseStmtBlock(a auditor) []ast.Stmt {
	var nodes []ast.Stmt

	for a.isNot(token.End) {
		nodes = append(nodes, expectStatement(a))
	}

	return nodes
}

// INLINE_STATEMENT := ASSIGNMENT | EXPR
func acceptInlineStatement(a auditor) (ast.Stmt, bool) {
	if a.is(token.Identifier) {
		return parseVariableStatement(a), true
	}

	if isNextLiteral(a) || isNextExprOpener(a) {
		expr := expectExpression(a)
		return operation(a, expr, 0), true
	}

	return nil, false
}

func expectStatement(a auditor) ast.Stmt {
	var (
		n  ast.Stmt
		ok bool
	)

	// TODO:
	// - for i, v in expr
	// - spell
	// - func
	// - proc

	if n, ok = acceptInlineStatement(a); ok {
		expectEndOfStmt(a)
		return n
	}

	if n, ok = acceptExpression(a); ok {
		expectEndOfStmt(a)
		return n
	}

	switch {
	case a.is(token.If):
		n = expectIf(a)

	case a.is(token.For):
		n = expectFor(a)

	case a.is(token.When):
		n = expectWhen(a)

	case a.is(token.Watch):
		n = expectWatch(a)

	default:
		panic(UnexpectedToken.Track("Expected statement"))
	}

	if n == nil {
		panic(err.New("Sanity check! Nil Node should never appear"))
	}

	expectEndOfStmt(a)
	return n
}

func parseVariableStatement(a auditor) ast.Stmt {
	first := a.expect(token.Identifier)

	if a.is(token.Comma) || a.is(token.Assign) {
		a.Putback(first)
		return expectAssignment(a)
	}

	a.Putback(first)
	return expectExpression(a)
}

func expectEndOfStmt(a auditor) {
	if !a.accept(token.Terminator) && !a.accept(token.Newline) {
		panic(a.unexpectedToken("Terminator or newline", a.Peek()))
	}
}
