package tokentest

import (
	"github.com/PaulioRandall/firefly-go/workflow/pos"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func Tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, pos.RawRangeForString(0, 0, 0, v))
}

type TokenGenerator func(token.TokenType, string) token.Token

func NewTokenGenerator() TokenGenerator {
	prev := pos.Range{}

	return func(tt token.TokenType, v string) token.Token {
		prev.ShiftString(v)
		return token.MakeToken(tt, v, prev)
	}
}
