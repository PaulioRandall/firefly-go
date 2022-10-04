package token

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/pos"
)

type Token struct {
	TokenType
	Value string
	Range pos.Range
}

func MakeToken(tt TokenType, val string, filePos pos.Range) Token {
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
