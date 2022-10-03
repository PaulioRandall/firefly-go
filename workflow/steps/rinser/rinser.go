// Package rinser removes obsolete tokens such as whitespace
package rinser

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var zero token.Token

type Input interface {
	More() bool
	Read() (token.Token, error)
}

type Output interface {
	Write(...token.Token) error
}

func Rinse(in Input, out Output) error {
	var prev token.Token

	for in.More() {
		tk, e := nextToken(in, prev)

		if errors.Is(e, inout.EOF) {
			return nil
		}

		if e != nil {
			return fmt.Errorf("Failed to rinse tokens: %w", e)
		}

		if tk == zero {
			continue
		}

		if e := out.Write(tk); e != nil {
			return fmt.Errorf("Failed to rinse tokens: %w", e)
		}

		prev = tk
	}

	return nil
}

func nextToken(in Input, prev token.Token) (token.Token, error) {
	switch tk, e := in.Read(); {
	case e != nil:
		return zero, e

	case tk.TokenType == token.Space:
		return zero, nil

	case tk.TokenType == token.Comment:
		return zero, nil

	case isEmptyLine(tk, prev):
		return zero, nil

	default:
		return tk, nil
	}
}

func isEmptyLine(tk, prev token.Token) bool {
	return tk.TokenType == token.Newline && prev.TokenType == token.Newline
}
