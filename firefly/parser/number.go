package parser

import (
	"strconv"

	"github.com/PaulioRandall/firefly-go/firefly/ast"
	"github.com/PaulioRandall/firefly-go/firefly/token"
)

func expectNumber(r lexReader) ast.Tree {

	lx := r.Peek()
	if lx.Token != token.TK_NUMBER {
		parsingPanic(nil, "Expected number, got '%s'", lx.Token.String())
	}

	return parseNumber(r.Read())
}

func parseNumber(num token.Lexeme) ast.Tree {

	n, e := strconv.ParseInt(num.Value, 10, 64)
	if e != nil {
		parsingPanic(e, "Unable to parse number '%s'", num.Value)
	}

	return ast.NumberTree{Value: n}
}
