package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func acceptStatements(a tokenAuditor) []ast.Stmt {
	var nodes []ast.Stmt

	for notEndOfBlock(a) {
		nodes = append(nodes, expectStatement(a))
	}

	return nodes
}

func expectStatement(a tokenAuditor) (n ast.Stmt) {
	switch {
	case accept(a, token.Identifier):
		n = expectVariableStatement(a, a.Prev())

	case isNext(a, token.If):
		n = parseIf(a)

	case isNext(a, token.When):
		n = expectWhen(a)

	default:
		panic(UnexpectedToken)
	}

	if n == nil {
		panic(err.New("Sanity check! Nil Node should never appear"))
	}

	expect(a, token.Terminator)
	return n
}

func expectVariableStatement(a tokenAuditor, first token.Token) ast.Stmt {
	if isNext(a, token.Comma) || isNext(a, token.Assign) {
		a.Putback(first)
		return expectAssignment(a)
	}

	panic(UnexpectedToken)
}
