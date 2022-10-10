// Package cleaner removes redundant tokens such as spaces and some newlines
package cleaner

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/models/token"
	"github.com/PaulioRandall/firefly-go/utilities/inout"
)

type TokenReader = inout.Reader[token.Token]
type TokenWriter = inout.Writer[token.Token]

func Clean(r TokenReader, w TokenWriter) error {
	e := inout.Process(r, w, processNext)
	if e != nil {
		return fmt.Errorf("Failed to clean tokens: %w", e)
	}
	return nil
}

func processNext(prev, curr, next token.Token) (token.Token, error) {
	var zero token.Token

	switch {
	case isRedundant(curr.TokenType):
		return zero, nil

	case curr.TokenType != token.Newline:
		return curr, nil

	case isEmptyLine(prev.TokenType):
		return zero, nil

	case isArithmetic(prev.TokenType):
		return zero, nil

	case isOpener(prev.TokenType):
		return zero, nil

	case isCloser(next.TokenType):
		return zero, nil

	default:
		return curr, nil
	}
}

func isRedundant(tt token.TokenType) bool {
	return tt == token.Space || tt == token.Comment
}

func isEmptyLine(tt token.TokenType) bool {
	return tt == token.Unknown || tt == token.Newline
}

func isArithmetic(tt token.TokenType) bool {
	switch tt {
	case token.Add, token.Sub, token.Mul, token.Div, token.Mod:
		return true
	default:
		return false
	}
}

func isOpener(tt token.TokenType) bool {
	switch tt {
	case token.ParenOpen, token.BraceOpen, token.BracketOpen:
		return true
	default:
		return false
	}
}

func isCloser(tt token.TokenType) bool {
	switch tt {
	case token.ParenClose, token.BraceClose, token.BracketClose:
		return true
	default:
		return false
	}
}
