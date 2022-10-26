package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectWatch(a auditor) ast.Watch {
	n := ast.Watch{}

	n.Keyword = a.expect(token.Watch)
	n.Variable = expectVariable(a)

	expectEndOfStmt(a)

	n.Body = acceptStatements(a)
	n.End = expectEndOfBlock(a)

	return n
}
