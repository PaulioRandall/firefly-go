package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectWhen(a tokenAuditor) ast.When {
	n := ast.When{}

	n.Keyword = expect(a, token.When)
	n.Subject = acceptExpression(a)

	expect(a, token.Terminator)

	n.Cases = acceptWhenCases(a)
	n.End = expect(a, token.End)

	return n
}

func acceptWhenCases(a tokenAuditor) []ast.WhenCase {
	var cases []ast.WhenCase

	for notEndOfBlock(a) {
		cases = append(cases, expectWhenCase(a))
		expect(a, token.Terminator)
	}

	return cases
}

func expectWhenCase(a tokenAuditor) ast.WhenCase {
	return ast.WhenCase{}
}
