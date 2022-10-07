// Package cleaner removes redundant tokens such as spaces and some newlines
package cleaner

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/process"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type Reader = inout.Reader[token.Token]
type Writer = inout.Writer[token.Token]

func Clean(r Reader, w Writer) error {
	e := process.Process(r, w, nextToken)
	if e != nil {
		return fmt.Errorf("Failed to clean tokens: %w", e)
	}
	return nil
}

func nextToken(prev, curr, next token.Token) (token.Token, bool, error) {
	var zero token.Token

	switch {
	case isRedundant(curr.TokenType):
		return zero, false, nil

	case curr.TokenType != token.Newline:
		return curr, true, nil

	case isEmptyLine(prev.TokenType):
		return zero, false, nil

	case isArithmetic(prev.TokenType):
		return zero, false, nil

	case isOpener(prev.TokenType):
		return zero, false, nil

	case isCloser(next.TokenType):
		return zero, false, nil

	default:
		return curr, true, nil
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
