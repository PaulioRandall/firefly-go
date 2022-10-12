package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectStatements(a *auditor) []ast.Stmt {
	var nodes []ast.Stmt

	for a.more() && !a.isNext(token.End) {
		nodes = append(nodes, expectStatement(a))
	}

	return nodes
}

func expectStatement(a *auditor) (n ast.Stmt) {
	switch {
	case a.accept(token.Identifier):
		n = expectVariableStatement(a, a.prev)

	case a.isNext(token.If):
		n = parseIf(a)

	default:
		panic(UnexpectedToken)
	}

	if n == nil {
		panic(err.New("Sanity check! Nil Node should never appear"))
	}

	a.expect(token.Terminator)
	return n
}

func expectVariableStatement(a *auditor, first token.Token) ast.Stmt {
	if a.isNext(token.Comma) || a.isNext(token.Assign) {
		a.putback(first)
		return expectAssignment(a)
	}

	panic(UnexpectedToken)
}
