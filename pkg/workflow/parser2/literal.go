package parser2

import (
	"strconv"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// == Number | String | True | False
func isLiteral(r BufReaderOfTokens) bool {
	tt := peekType(r)
	return tt == token.Number ||
		tt == token.String ||
		tt == token.True ||
		tt == token.False
}

// LITERAL := NUMBER | STRING | True | False
func parseLiteral(r BufReaderOfTokens) ast.Literal {
	switch peekType(r) {
	case token.Number:
		return parseNumber(r)
	case token.String:
		return parseString(r)
	case token.True, token.False:
		return parseBool(r)
	default:
		panic(ErrParsing.Track("Expected literal"))
	}
}

// NUMBER := Number
func parseNumber(r BufReaderOfTokens) ast.Literal {
	v := expectType(r, token.Number).Value
	num, e := strconv.ParseFloat(v, 64)

	if e != nil {
		panic(ErrParsing.Track("Unable to parse number"))
	}

	return ast.Literal{
		Value: num,
	}
}

// STRING := String
func parseString(r BufReaderOfTokens) ast.Literal {
	str := expectType(r, token.String).Value
	str = str[1 : len(str)-1] // Slice off delimiters
	return ast.Literal{
		Value: str,
	}
}

// BOOL := Bool
func parseBool(r BufReaderOfTokens) ast.Literal {
	s := readToken(r).Value // expectType(r, token.Bool)
	b, e := strconv.ParseBool(s)

	if e != nil {
		panic(ErrParsing.Track("Unable to parse bool"))
	}

	return ast.Literal{
		Value: b,
	}
}
