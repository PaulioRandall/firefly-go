// Package terminator converts newlines in to terminators
package terminator

import (
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var zeroTk token.Token

type TokenReader interface {
	More() bool
	Peek() (token.Token, error)
	Read() (token.Token, error)
}

type TokenWriter interface {
	Write(token.Token) error
}

func Terminate(r TokenReader, w TokenWriter) error {
	for r.More() {
		curr, e := r.Read()
		if e != nil {
			return e // TODO: Wrap error
		}

		if curr.TokenType == token.Newline {
			curr.TokenType = token.Terminator
		}

		if e = w.Write(curr); e != nil {
			return e // TODO: Wrap error
		}
	}

	return nil
}
