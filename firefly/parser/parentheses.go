package parser

import (
	"github.com/PaulioRandall/firefly-go/firefly/ast"
	"github.com/PaulioRandall/firefly-go/firefly/token"
)

func parseParenExpr(r lexReader, opener token.Lexeme) ast.Tree {

	if !r.More() {
		parsingPanic(nil, "Expected expression after opening parenthesis '('")
	}

	tr := expectExpr(r, 0)
	expectParenClose(r)

	return tr
}

func expectParenClose(r lexReader) {
	tk := r.Read().Token
	if tk != token.TK_PAREN_CLOSE {
		parsingPanic(nil, "Expected closing parenthesis but got '%s'", tk.String())
	}
}
