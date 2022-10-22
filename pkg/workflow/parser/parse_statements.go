package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func acceptStatements(r PosReaderOfTokens) []ast.Stmt {
	var nodes []ast.Stmt

	for isNotEndOfBlock(r) {
		nodes = append(nodes, expectStatement(r))
	}

	return nodes
}

func expectStatement(r PosReaderOfTokens) (n ast.Stmt) {
	switch {
	case accept(r, token.Identifier):
		n = expectVariableStatement(r, r.Prev())

	case is(r, token.If):
		n = parseIf(r)

	case is(r, token.When):
		n = expectWhen(r)

	case match(r, token.IsLiteral):
		n = expectLiteral(r)

	default:
		panic(UnexpectedToken)
	}

	if n == nil {
		panic(err.New("Sanity check! Nil Node should never appear"))
	}

	expect(r, token.Terminator)
	return n
}

func expectVariableStatement(r PosReaderOfTokens, first token.Token) ast.Stmt {
	if is(r, token.Comma) || is(r, token.Assign) {
		r.Putback(first)
		return expectAssignment(r)
	}

	panic(UnexpectedToken)
}
