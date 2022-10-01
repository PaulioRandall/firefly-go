package tokentest

import (
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type TokenGenerator func(token.TokenType, string) token.Token

func NewTokenGenerator() TokenGenerator {
	prev := token.Range{}

	return func(tt token.TokenType, v string) token.Token {
		prev.IncString(v)
		return token.MakeToken(tt, v, prev)
	}
}

func Tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, token.Range{})
}

func InlineRange(offset, line, col, length int) token.Range {
	return token.Range{
		From: token.Pos{
			Offset: offset,
			Line:   line,
			Col:    col,
		},
		To: token.Pos{
			Offset: offset + length,
			Line:   line,
			Col:    col + length,
		},
	}
}
