// Package aligner creates easy to parse comma separated tokens
package aligner

import (
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type TokenReader interface {
	More() bool
	Read() (token.Token, error)
	Peek() (token.Token, error)
}

type TokenWriter interface {
	Write(token.Token) error
}

func Align(r TokenReader, w TokenWriter) error {
	var prev token.Token

	for r.More() {
		curr, e := r.Read()
		if e != nil {
			return e // TODO: Wrap error using current/previous token position
		}

		if remove, e := removeToken(r, prev, curr); e != nil {
			return e // TODO: Wrap error using current/previous token position
		} else if remove {
			continue
		}

		if e := w.Write(curr); e != nil {
			return e // TODO: Wrap error using current/previous token position
		}

		prev = curr
	}

	return nil
}

func removeToken(r TokenReader, prev, curr token.Token) (bool, error) {
	if isCSVTerminator(prev, curr) {
		return true, nil
	}

	if !r.More() {
		return false, nil
	}

	next, e := r.Peek()
	if e != nil {
		return false, e // TODO: Wrap error using current/previous token position
	}

	return isTrailingCSVComma(curr, next), nil
}

func isCSVTerminator(prev, curr token.Token) bool {
	return prev.TokenType == token.Comma && curr.TokenType == token.Terminator
}

func isTrailingCSVComma(curr, next token.Token) bool {
	return curr.TokenType == token.Comma && isCloser(next.TokenType)
}

func isCloser(tt token.TokenType) bool {
	switch tt {
	case token.ParenClose, token.BraceClose, token.BracketClose:
		return true
	default:
		return false
	}
}
