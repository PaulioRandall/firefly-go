// Package aligner creates easy to parse comma separated tokens
package aligner

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ReaderOfTokens = inout.Reader[token.Token]
type WriterOfTokens = inout.Writer[token.Token]

var ErrAligning = err.Trackable("Aligning failed")

func Align(r ReaderOfTokens, w WriterOfTokens) error {
	e := inout.Stream(r, w, processNext)
	if e != nil {
		return ErrAligning.Track(e, "Aligner failed to align tokens")
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
