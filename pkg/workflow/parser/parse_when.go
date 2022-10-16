package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectWhen(r BufReaderOfTokens) ast.When {
	n := ast.When{}

	n.Keyword = expect(r, token.When)
	n.Subject = acceptExpression(r)

	expect(r, token.Terminator)

	n.Cases = acceptWhenCases(r)
	n.End = expect(r, token.End)

	return n
}

func acceptWhenCases(r BufReaderOfTokens) []ast.WhenCase {
	var cases []ast.WhenCase

	for notEndOfBlock(r) {
		cases = append(cases, expectWhenCase(r))
		expect(r, token.Terminator)
	}

	return cases
}

func expectWhenCase(r BufReaderOfTokens) ast.WhenCase {
	return ast.WhenCase{}
}
