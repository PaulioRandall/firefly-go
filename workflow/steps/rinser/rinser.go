// Package rinser removes whitespace and syntactic sugar from a list of tokens
package rinser

import (
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var (
	zeroToken token.Token
)

type TokenReader interface {
	More() bool
	Read() token.Token
	Peek() token.Token
}

type RinseFunc func() (tk token.Token, e error)

func New(tr TokenReader) RinseFunc {
	var prev token.Token

	return func() (token.Token, error) {
		for tr.More() {
			tk, removed, e := nextToken(tr)
			if e != nil {
				return zeroToken, err.AfterToken(prev, e, "Failed to rinse token")
			}

			if !removed {
				prev = tk
				return tk, nil
			}
		}

		return zeroToken, err.EOF
	}
}

func nextToken(tr TokenReader) (token.Token, bool, error) {
	return zeroToken, false, nil
}
