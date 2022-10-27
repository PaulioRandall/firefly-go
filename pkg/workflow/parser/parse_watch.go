package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectWatch(a auditor) ast.Watch {
	n := ast.Watch{}

	n.Keyword = a.expect(token.Watch)
	n.Variable, _ = acceptVariable(a)

	expectEndOfStmt(a)

	n.Body = parseStmtBlock(a)
	n.End = a.expect(token.End)

	return n
}
