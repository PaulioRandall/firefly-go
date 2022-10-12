// Package aligner creates easy to parse comma separated tokens
package aligner

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type TokenReader = inout.Reader[token.Token]
type TokenWriter = inout.Writer[token.Token]

func Align(r TokenReader, w TokenWriter) error {
	e := inout.Stream(r, w, processNext)
	if e != nil {
		return err.Wrap(e, "Aligner failed to align tokens")
	}
	return nil
}

func processNext(prev, curr, next token.Token) (token.Token, error) {
	var zero token.Token

	if isCommaBeforeTerminator(prev, curr) {
		return zero, nil
	}

	if isCommaBeforeClosingParentheses(curr, next) {
		return zero, nil
	}

	return curr, nil
}

func isCommaBeforeTerminator(prev, curr token.Token) bool {
	return prev.TokenType == token.Comma && curr.TokenType == token.Terminator
}

func isCommaBeforeClosingParentheses(curr, next token.Token) bool {
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
