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

func acceptInlineStatement(a auditor) ast.Stmt {
	switch {
	case a.accept(token.Identifier):
		return expectVariableStatement(a, a.Prev())

	case a.is(token.BracketOpen), a.is(token.BraceOpen):
		return expectExpression(a)

	case a.match(token.IsLiteral), a.is(token.ParenOpen):
		expr := expectExpression(a)
		return operation(a, expr, 0)

	default:
		return nil
	}
}

func expectStatement(a auditor) (n ast.Stmt) {

	// TODO:
	// - for i, v in expr
	// - spell
	// - func
	// - proc

	switch {
	case a.accept(token.Identifier):
		n = expectVariableStatement(a, a.Prev())

	case a.is(token.If):
		n = expectIf(a)

	case a.is(token.For):
		n = expectFor(a)

	case a.is(token.When):
		n = expectWhen(a)

	case a.is(token.Watch):
		n = expectWatch(a)

	case a.is(token.BracketOpen), a.is(token.BraceOpen):
		n = expectExpression(a)

	case a.match(token.IsLiteral), a.is(token.ParenOpen):
		expr := expectExpression(a)
		n = operation(a, expr, 0)

	default:
		panic(UnexpectedToken.Track("Expected statement"))
	}

	if n == nil {
		panic(err.New("Sanity check! Nil Node should never appear"))
	}

	expectEndOfStmt(a)
	return n
}

func expectVariableStatement(a auditor, first token.Token) ast.Stmt {
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
