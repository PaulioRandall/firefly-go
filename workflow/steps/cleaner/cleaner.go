// Package cleaner removes compiler redundant tokens such as whitespace
package rinser

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var zero token.Token

type TokenReader interface {
	More() bool
	Read() (token.Token, error)
}

type TokenWriter interface {
	Write(token.Token) error
}

func Clean(r TokenReader, w TokenWriter) error {
	var prev token.Token

	for r.More() {
		tk, e := nextToken(r, prev)

		if errors.Is(e, inout.EOF) {
			return nil
		}

		if e != nil {
			return fmt.Errorf("Failed to clean tokens: %w", e)
		}

		if tk == zero {
			continue
		}

		if e := w.Write(tk); e != nil {
			return fmt.Errorf("Failed to clean tokens: %w", e)
		}

		prev = tk
	}

	return nil
}

func nextToken(r TokenReader, prev token.Token) (token.Token, error) {
	switch tk, e := r.Read(); {
	case e != nil:
		return zero, e

	case isRedundant(tk.TokenType):
		return zero, nil

	case isEmptyLine(tk.TokenType, prev.TokenType):
		return zero, nil

	default:
		return tk, nil
	}
}

func isRedundant(tt token.TokenType) bool {
	switch tt {
	case token.Space, token.Comment:
		return true
	default:
		return false
	}
}

func isEmptyLine(curr, prev token.TokenType) bool {
	if curr != token.Newline {
		return false
	}

	return prev == token.Unknown || prev == token.Newline
}
