package parser2

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// == Identifier Comma
// == Identifier Assign
func isAssignment(r BufReaderOfTokens) bool {
	if peekType(r) != token.Identifier {
		return false
	}

	ident := readToken(r)
	defer r.Putback(ident)

	tt := peekType(r)
	return tt == token.Comma || tt == token.Assign
}

// ASSIGNMENT := VARIABLES Assign EXPRESSIONS
func parseAssignment(r BufReaderOfTokens) ast.Assign {
	n := ast.Assign{}

	n.Dst = parseVariables(r)
	expectType(r, token.Assign)
	n.Src = parseExpressions(r)

	return n
}
