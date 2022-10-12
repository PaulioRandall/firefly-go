package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
)

func parseIf(a *auditor.Auditor) ast.If {
	n := ast.If{}

	n.Keyword = a.Expect(token.If)
	n.Condition = expectExpression(a)

	a.Expect(token.Terminator)

	n.Body = expectStatements(a)
	n.End = a.Expect(token.End)

	return n
}
