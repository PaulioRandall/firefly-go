// Package aligner creates easy to parse comma separated tokens
package aligner

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type TokenReader = inout.Reader[token.Token]
type TokenWriter = inout.Writer[token.Token]

func Align(r TokenReader, w TokenWriter) error {
	e := inout.Process(r, w, processNext)
	if e != nil {
		return fmt.Errorf("Failed to align tokens: %w", e)
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
