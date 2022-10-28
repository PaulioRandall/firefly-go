package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrBadWhen              = err.Trackable("Failed to parse when statement")
	ErrBadWhenCase          = err.Trackable("Failed to parse when case")
	ErrBadWhenCaseCondition = err.Trackable("Failed to parse when case condition")
)

func acceptWhen(a auditor) (ast.When, bool) {
	if a.isNot(token.When) {
		return ast.When{}, false
	}
	return expectWhen(a), true
}

func expectWhen(a auditor) ast.When {
	defer wrapPanic(func(e error) error {
		return ErrBadWhen.Wrap(e, "Bad when statement syntax")
	})

	n := ast.When{}

	n.Keyword = a.expect(token.When)
	n.Subject, _ = acceptExpression(a)

	expectEndOfStmt(a)

	n.Cases = parseWhenBlock(a)
	n.End = a.expect(token.End)

	return n
}

func parseWhenBlock(a auditor) []ast.WhenCase {
	var cases []ast.WhenCase

	for a.isNot(token.End) {
		cases = append(cases, expectWhenCase(a))
	}

	return cases
}

func expectWhenCase(a auditor) ast.WhenCase {
	defer wrapPanic(func(e error) error {
		return ErrBadWhenCase.Wrap(e, "Bad when case")
	})

	condition := expectWhenCaseCondition(a)
	a.expect(token.Colon)

	var body ast.Stmt
	if a.is(token.Terminator) {
		expectEndOfStmt(a)
	} else {
		body = expectStatement(a)
	}

	return ast.WhenCase{
		Condition: condition,
		Statement: body,
	}
}

func expectWhenCaseCondition(a auditor) ast.Expr {
	defer wrapPanic(func(e error) error {
		return ErrBadWhenCaseCondition.Wrap(e, "Bad when case condition syntax")
	})

	if a.isNot(token.Is) {
		return expectExpression(a)
	}

	return ast.Is{
		Keyword: a.expect(token.Is),
		Expr:    expectExpression(a),
	}
}
