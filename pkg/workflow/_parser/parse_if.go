package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrBadIfStmt = err.Trackable("Failed to parse if statement")
)

func acceptIf(a auditor) (ast.If, bool) {
	if a.isNot(token.If) {
		return ast.If{}, false
	}
	return expectIf(a), true
}

func expectIf(a auditor) ast.If {
	defer wrapPanic(func(e error) error {
		return ErrBadIfStmt.Wrap(e, "Bad if statement syntax")
	})

	n := ast.If{}

	n.Keyword = a.expect(token.If)
	n.Condition = expectExpression(a)

	expectEndOfStmt(a)

	n.Body = parseStmtBlock(a)
	n.End = parseEndOfBlock(a)

	return n
}
