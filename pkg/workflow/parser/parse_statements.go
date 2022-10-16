package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func acceptStatements(r PosReaderOfTokens) []ast.Stmt {
	var nodes []ast.Stmt

	for notEndOfBlock(r) {
		nodes = append(nodes, expectStatement(r))
	}

	return nodes
}

func expectStatement(r PosReaderOfTokens) (n ast.Stmt) {
	switch {
	case accept(r, token.Identifier):
		n = expectVariableStatement(r, r.Prev())

	case isNext(r, token.If):
		n = parseIf(r)

	case isNext(r, token.When):
		n = expectWhen(r)

	case doesNextMatch(r, token.IsLiteral):
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
	if isNext(r, token.Comma) || isNext(r, token.Assign) {
		r.Putback(first)
		return expectAssignment(r)
	}

	panic(UnexpectedToken)
}
