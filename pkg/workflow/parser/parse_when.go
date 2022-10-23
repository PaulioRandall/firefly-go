package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectWhen(a auditor) ast.When {
	n := ast.When{}

	n.Keyword = a.expect(token.When)
	n.Subject = acceptExpression(a)

	expectTerminator(a)

	n.Cases = acceptWhenCases(a)
	n.End = expectEndOfBlock(a)

	return n
}

func acceptWhenCases(a auditor) []ast.WhenCase {
	var cases []ast.WhenCase

	for isNotEndOfBlock(a) {
		cases = append(cases, expectWhenCase(a))
	}

	return cases
}

func expectWhenCase(a auditor) ast.WhenCase {

	condition := expectExpression(a)
	a.expect(token.Colon)

	var body ast.Stmt
	if a.is(token.Terminator) {
		expectTerminator(a)
	} else {
		body = expectStatement(a)
	}

	return ast.WhenCase{
		Condition: condition,
		Statement: body,
	}
}
