package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrBadWatchStmt = err.Trackable("Failed to parse watch statement")
)

func acceptWatch(a auditor) (ast.Watch, bool) {
	if a.isNot(token.Watch) {
		return ast.Watch{}, false
	}
	return expectWatch(a), true
}

func expectWatch(a auditor) ast.Watch {
	defer wrapPanic(func(e error) error {
		return ErrBadWatchStmt.Wrap(e, "Bad watch statement")
	})

	n := ast.Watch{}

	n.Keyword = a.expect(token.Watch)
	n.Variable = expectVariable(a)

	expectEndOfStmt(a)

	n.Body = parseStmtBlock(a)
	n.End = parseEndOfBlock(a)

	return n
}
