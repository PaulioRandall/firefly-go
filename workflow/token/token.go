package token

import (
	"fmt"
)

type Token struct {
	Range
	Type  TokenType
	Value string
}

func MakeToken(tt TokenType, val string, filePos Range) Token {
	return Token{
		Type:  tt,
		Value: val,
		Range: filePos,
	}
}

func (tk Token) String() string {
	return fmt.Sprintf("%s %q %s", tk.Type, tk.Value, tk.Range)
}

type TokenGenerator func(TokenType, string) Token

func NewTokenGenerator() TokenGenerator {
	prev := Range{}

	return func(tt TokenType, v string) Token {
		prev.IncString(v)
		return MakeToken(tt, v, prev)
	}
}
