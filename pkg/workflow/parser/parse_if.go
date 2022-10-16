package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
)

func parseIf(a *auditor.Auditor) ast.If {
	n := ast.If{}

	n.Keyword = expect(a, token.If)
	n.Condition = expectExpression(a)

	expect(a, token.Terminator)

	n.Body = acceptStatements(a)
	n.End = expect(a, token.End)

	return n
}
