// Package cleaner removes redundant tokens such as spaces and some newlines.
package cleaner2

import (
	"github.com/PaulioRandall/go-trackerr"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

var zeroToken token.Token

type ReaderOfTokens interface {
	More() bool
	Read() token.Token
}

type WriterOfTokens interface {
	Write(token.Token) error
}

var ErrCleaning = trackerr.Checkpoint("Cleaning failed")

func Clean(r ReaderOfTokens, w WriterOfTokens) error {
	var prev, curr, next token.Token

	if !r.More() {
		return nil
	}

	for next = r.Read(); next != zeroToken; {
		if curr != zeroToken {
			prev = curr
		}

		curr = next

		if r.More() {
			next = r.Read()
		} else {
			next = zeroToken
		}

		curr = clean(prev, curr, next)

		if curr == zeroToken {
			continue
		}

		if e := w.Write(curr); e != nil {
			return ErrCleaning.CausedBy(e, "Failed to write token")
		}
	}

	return nil
}

func clean(prev, curr, next token.Token) token.Token {
	switch {
	case isRedundant(curr.TokenType):
		return zeroToken

	case curr.TokenType != token.Newline:
		return curr

		// curr == Newline from here on
	case isEmptyLine(prev.TokenType):
		return zeroToken

	case isBinaryOperator(prev.TokenType):
		return zeroToken

	case isOpener(prev.TokenType):
		return zeroToken

	case isCloser(next.TokenType):
		return zeroToken

	default:
		return curr
	}
}

func isRedundant(tt token.TokenType) bool {
	return tt == token.Space || tt == token.Comment
}

func isEmptyLine(tt token.TokenType) bool {
	return tt == token.Unknown || tt == token.Newline
}

func isBinaryOperator(tt token.TokenType) bool {
	switch tt {
	case token.Add, token.Sub, token.Mul, token.Div, token.Mod:
	case token.Lt, token.Gt, token.Lte, token.Gte, token.Equ, token.Neq:
	case token.And, token.Or:
	default:
		return false
	}

	return true
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
