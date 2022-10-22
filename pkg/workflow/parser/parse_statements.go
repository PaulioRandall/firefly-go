package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func acceptStatements(a auditor) []ast.Stmt {
	var nodes []ast.Stmt

	for isNotEndOfBlock(a) {
		nodes = append(nodes, expectStatement(a))
	}

	return nodes
}

func expectStatement(a auditor) (n ast.Stmt) {
	switch {
	case a.accept(token.Identifier):
		n = expectVariableStatement(a, a.Prev())

	case a.is(token.If):
		n = parseIf(a)

	case a.is(token.When):
		n = expectWhen(a)

	case a.match(token.IsLiteral):
		n = expectExpression(a)

	default:
		panic(UnexpectedToken)
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

	panic(UnexpectedToken)
}
