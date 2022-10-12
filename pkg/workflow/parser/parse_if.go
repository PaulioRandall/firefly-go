package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func parseIf(a *auditor) ast.If {
	n := ast.If{}

	n.Keyword = a.expect(token.If)
	n.Condition = expectExpression(a)

	a.expect(token.Terminator)

	n.Body = expectStatements(a)
	n.End = a.expectWith(MissingEnd, token.End)

	return n
}
