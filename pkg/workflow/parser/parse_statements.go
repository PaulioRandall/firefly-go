package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
)

func expectStatements(a *auditor.Auditor) []ast.Stmt {
	var nodes []ast.Stmt

	for a.More() && !a.IsNext(token.End) {
		nodes = append(nodes, expectStatement(a))
	}

	return nodes
}

func expectStatement(a *auditor.Auditor) (n ast.Stmt) {
	switch {
	case a.Accept(token.Identifier):
		n = expectVariableStatement(a, a.Prev())

	case a.IsNext(token.If):
		n = parseIf(a)

	default:
		panic(auditor.UnexpectedToken)
	}

	if n == nil {
		panic(err.New("Sanity check! Nil Node should never appear"))
	}

	a.Expect(token.Terminator)
	return n
}

func expectVariableStatement(a *auditor.Auditor, first token.Token) ast.Stmt {
	if a.IsNext(token.Comma) || a.IsNext(token.Assign) {
		a.Putback(first)
		return expectAssignment(a)
	}

	panic(auditor.UnexpectedToken)
}
