// Package rinser removes whitespace and syntactic sugar from a list of tokens
package rinser

import (
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var (
	zero token.Token
)

type TokenReader interface {
	More() bool
	Read() token.Token
	Peek() token.Token
}

type RinseNext func() (tk token.Token, e error)

func New(tr TokenReader) RinseNext {
	return func() (token.Token, error) {
		for tr.More() {
			if tk := nextToken(tr); tk != zero {
				return tk, nil
			}
		}

		return zero, err.EOF
	}
}

func nextToken(tr TokenReader) token.Token {
	switch tk := tr.Read(); {
	case tk.Type == token.Space:
		return zero
	default:
		return tk
	}
}
