// Package cleaner removes redundant tokens such as spaces and some newlines
package cleaner

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var zero token.Token

type TokenReader interface {
	More() bool
	Peek() (token.Token, error)
	Read() (token.Token, error)
}

type TokenWriter interface {
	Write(token.Token) error
}

func Clean(r TokenReader, w TokenWriter) error {
	var prev token.Token

	for r.More() {
		curr, e := nextToken(r, prev)

		if errors.Is(e, inout.EOF) {
			return nil
		}

		if e != nil {
			return fmt.Errorf("Failed to clean tokens: %w", e)
		}

		if curr == zero {
			continue
		}

		if e := w.Write(curr); e != nil {
			return fmt.Errorf("Failed to clean tokens: %w", e)
		}

		prev = curr
	}

	return nil
}

func nextToken(r TokenReader, prev token.Token) (token.Token, error) {
	switch curr, next, e := readPeek(r); {
	case e != nil:
		return zero, e

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

func readPeek(r TokenReader) (curr, next token.Token, e error) {
	if curr, e = r.Read(); e != nil {
		return
	}

	if r.More() {
		next, e = r.Peek()
	}

	return
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
