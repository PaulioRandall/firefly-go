package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectIf(a auditor) ast.If {
	n := ast.If{}

	n.Keyword = a.expect(token.If)
	n.Condition = expectExpression(a)

	expectEndOfStmt(a)

	n.Body = acceptStatements(a)
	n.End = expectEndOfBlock(a)

	return n
}
