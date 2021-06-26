package parser

import (
	"strconv"

	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func expectNumber(r lexReader) ast.Node {

	lx := r.Read()
	if lx.Token != token.TokenNumber {
		parsingPanic(nil, "Expected number, got '%s'", lx.Token.String())
	}

	return parseNumber(lx)
}

func parseNumber(num token.Lexeme) ast.Node {

	n, e := strconv.ParseInt(num.Value, 10, 64)
	if e != nil {
		parsingPanic(e, "Unable to parse number '%s'", num.Value)
	}

	return ast.NumberNode{Value: n}
}
