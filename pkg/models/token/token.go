package token

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
)

type Token struct {
	TokenType
	Value string
	From  pos.Pos
	To    pos.Pos
}

func MakeToken(tt TokenType, v string, from, to pos.Pos) Token {
	return Token{
		TokenType: tt,
		Value:     v,
		From:      from,
		To:        to,
	}
}

func MakeTokenAt(tt TokenType, v string, from pos.Pos) Token {
	return Token{
		TokenType: tt,
		Value:     v,
		From:      from,
		To:        pos.EndOf(from, v),
	}
}

func (tk Token) Where() (pos.Pos, pos.Pos) {
	return tk.From, tk.To
}

func (tk Token) Debug() string {
	return fmt.Sprintf("%s %q %s", tk.TokenType.String(), tk.Value, pos.WhereString(tk))
}

func (tk Token) String() string {
	return fmt.Sprintf("%s %q", tk.TokenType.String(), tk.Value)
}
