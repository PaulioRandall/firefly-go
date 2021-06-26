package parser

import (
	"strconv"

	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func expectNumber(lr token.LexemeReader) (ast.Node, error) {

	lx, e := lr.Read()
	if e != nil {
		return nil, e
	}

	if lx.Token != token.TokenNumber {
		return nil, newError("Expected number, got '%s'", lx.Token.String())
	}
	
	return parseNumber(lx)
}

func parseNumber(num token.Lexeme) (ast.Node, error) {
	n, e := strconv.ParseInt(num.Value, 10, 64)
	if e != nil {
		wrapThenPanic(e, "Unable to parse number '%s'", num.Value)
	}
	return ast.NumberNode{Value: n}, nil
}
