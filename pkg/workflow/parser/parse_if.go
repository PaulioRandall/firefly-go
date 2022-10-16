package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func parseIf(r PosReaderOfTokens) ast.If {
	n := ast.If{}

	n.Keyword = expect(r, token.If)
	n.Condition = expectExpression(r)

	expect(r, token.Terminator)

	n.Body = acceptStatements(r)
	n.End = expect(r, token.End)

	return n
}
