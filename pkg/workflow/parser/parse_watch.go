package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrBadWatch = err.Trackable("Failed to parse watch statement")
)

func expectWatch(a auditor) ast.Watch {
	defer wrapPanic(func(e error) error {
		return ErrBadWatch.Wrap(e, "Bad watch statement syntax")
	})

	n := ast.Watch{}

	n.Keyword = a.expect(token.Watch)
	n.Variable, _ = acceptVariable(a)

	expectEndOfStmt(a)

	n.Body = parseStmtBlock(a)
	n.End = a.expect(token.End)

	return n
}
