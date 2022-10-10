// Package terminator converts all newlines in to terminators
package terminator

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/models/token"
	"github.com/PaulioRandall/firefly-go/utilities/inout"
)

type TokenReader = inout.Reader[token.Token]
type TokenWriter = inout.Writer[token.Token]

func Terminate(r TokenReader, w TokenWriter) error {
	e := inout.Process(r, w, processNext)
	if e != nil {
		return fmt.Errorf("Failed to convert newlines to terminators: %w", e)
	}
	return nil
}

func processNext(prev, curr, next token.Token) (token.Token, error) {
	if curr.TokenType == token.Newline {
		curr.TokenType = token.Terminator
	}
	return curr, nil
}
