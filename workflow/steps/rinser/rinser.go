// Package rinser removes obsolete tokens such as whitespace
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
	var prev token.Token

	return func() (token.Token, error) {
		for tr.More() {
			if tk := nextToken(tr, prev); tk != zero {
				prev = tk
				return tk, nil
			}
		}

		return zero, err.EOF
	}
}

func nextToken(tr TokenReader, prev token.Token) token.Token {
	switch tk := tr.Read(); {
	case tk.Type == token.Space:
		return zero

	case tk.Type == token.Comment:
		return zero

	case isEmptyLine(tk, prev):
		return zero

	default:
		return tk
	}
}

func isEmptyLine(tk, prev token.Token) bool {
	return tk.Type == token.Newline && prev.Type == token.Newline
}
