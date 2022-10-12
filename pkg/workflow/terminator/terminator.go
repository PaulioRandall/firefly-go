// Package terminator converts all newlines in to terminators
package terminator

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type TokenReader = inout.Reader[token.Token]
type TokenWriter = inout.Writer[token.Token]

func Terminate(r TokenReader, w TokenWriter) error {
	e := inout.Stream(r, w, processNext)
	if e != nil {
		return err.Wrap(e, "Terminator failed to convert newlines to terminators")
	}
	return nil
}

func processNext(prev, curr, next token.Token) (token.Token, error) {
	if curr.TokenType == token.Newline {
		curr.TokenType = token.Terminator
	}
	return curr, nil
}
