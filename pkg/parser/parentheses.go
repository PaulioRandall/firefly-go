package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func parseParenExpr(lr token.LexemeReader, opener token.Lexeme) (ast.Node, error) {

	if !lr.More() {
		return nil, newError("Expected expression after opening parenthesis '('")
	}

	n, e := expectExpr(lr, 0)
	if e != nil {
		return nil, e
	}

	e = expectParenClose(lr)
	if e != nil {
		return nil, e
	}

	return n, nil
}

func expectParenClose(lr token.LexemeReader) error {

	lx, e := lr.Read()
	if e != nil {
		return e
	}

	if lx.Token != token.TokenParenClose {
		tk := lx.Token.String()
		return newError("Expected closing parenthesis but got '%s'", tk)
	}

	return nil
}
