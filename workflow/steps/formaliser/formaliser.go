// Package formaliser modifies a token list for easier parsing
package formaliser

import (
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type tokenList []token.Token

func (tl *tokenList) append(tk token.Token) {
	*tl = append(*tl, tk)
}

func (tl *tokenList) last() token.Token {
	if i := tl.indexOfLast(); i > -1 {
		return (*tl)[i]
	}
	return token.Token{}
}

func (tl *tokenList) indexOfLast() int {
	return len(*tl) - 1
}

var zeroTk token.Token

type TokenReader interface {
	More() bool
	Read() (token.Token, error)
}

type TokenWriter interface {
	Write(token.Token) error
}

func Formalise(r TokenReader, w TokenWriter) error {
	var (
		prev token.Token
		curr token.Token
		e    error
	)

	for r.More() {
		if curr, e = r.Read(); e != nil {
			return e // TODO: Wrap error
		}

		if curr = formalise(r, prev, curr); curr == zeroTk {
			continue
		}

		if e = w.Write(curr); e != nil {
			return e // TODO: Wrap error
		}

		prev = curr
	}

	return nil
}

func formalise(r TokenReader, prev, curr token.Token) token.Token {
	if curr.TokenType != token.Newline {
		return curr
	}

	if preventsNewlineTermination(prev) {
		return zeroTk
	}

	curr.TokenType = token.Terminator
	return curr
}

func preventsNewlineTermination(tk token.Token) bool {
	switch tk.TokenType {
	case token.Add, token.Sub, token.Mul, token.Div, token.Mod:
		return true
	default:
		return false
	}
}
