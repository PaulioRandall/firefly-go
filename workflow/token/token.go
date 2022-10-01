package token

import (
	"fmt"
)

type Token struct {
	TokenType
	Value string
	Range
}

func MakeToken(tt TokenType, val string, filePos Range) Token {
	return Token{
		TokenType: tt,
		Value:     val,
		Range:     filePos,
	}
}

func (tk Token) Debug() string {
	return fmt.Sprintf("%s %q %s", tk.TokenType.String(), tk.Value, tk.Range)
}

func (tk Token) String() string {
	return fmt.Sprintf("%s: %q", tk.TokenType.String(), tk.Value)
}
