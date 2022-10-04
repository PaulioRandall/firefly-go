// Package terminator removes newlines or converts them to terminators
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
	var (
		prev token.Token
		curr token.Token
		next token.Token
		e    error
	)

	for r.More() {
		if curr, e = r.Read(); e != nil {
			return e // TODO: Wrap error
		}

		if r.More() {
			if next, e = r.Peek(); e != nil {
				return e // TODO: Wrap error
			}
		}

		if curr = terminate(prev, curr, next); curr == zeroTk {
			continue
		}

		if e = w.Write(curr); e != nil {
			return e // TODO: Wrap error
		}

		prev = curr
	}

	return nil
}

func terminate(prev, curr, next token.Token) token.Token {
	switch {
	case curr.TokenType != token.Newline:
		return curr
	case isArithmetic(prev.TokenType):
		return zeroTk
	case isOpener(prev.TokenType):
		return zeroTk
	case isCloser(next.TokenType):
		return zeroTk
	default:
		curr.TokenType = token.Terminator
		return curr
	}
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
