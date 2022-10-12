package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
)

func expectWhen(a *auditor.Auditor) ast.When {
	n := ast.When{}

	n.Keyword = a.Expect(token.When)
	n.Subject = acceptExpression(a)

	a.Expect(token.Terminator)

	n.Cases = acceptWhenCases(a)
	n.End = a.Expect(token.End)

	return n
}

func acceptWhenCases(a *auditor.Auditor) []ast.WhenCase {
	var cases []ast.WhenCase

	for notEndOfBlock(a) {
		cases = append(cases, expectWhenCase(a))
		a.Expect(token.Terminator)
	}

	return cases
}

func expectWhenCase(a *auditor.Auditor) ast.WhenCase {
	return ast.WhenCase{}
}
