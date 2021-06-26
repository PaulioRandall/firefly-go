package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func parseParenExpr(r lexReader, opener token.Lexeme) ast.Node {

	if !r.More() {
		parsingPanic(nil, "Expected expression after opening parenthesis '('")
	}

	n := expectExpr(r, 0)
	expectParenClose(r)

	return n
}

func expectParenClose(r lexReader) {
	tk := r.Read().Token
	if tk != token.TK_PAREN_CLOSE {
		parsingPanic(nil, "Expected closing parenthesis but got '%s'", tk.String())
	}
}
