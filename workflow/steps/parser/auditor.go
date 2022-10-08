package parser

import (
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var zero token.Token

type TokenReader = inout.Reader[token.Token]

type auditor struct {
	TokenReader
	next token.Token // TODO: impl but only used for put backs
	prev token.Token
}

func (a auditor) getPrev() token.Token {
	return a.prev
}

/*
func (a *auditor) putBack() {
	if a.prev == zero {
		e := err.AtToken(a.prev, nil, "Can't put back that which has not been read")
		panic(e)
	}

	a.next, a.prev = a.prev, zero
}
*/
func (a *auditor) isNext(want token.TokenType) bool {
	return a.doesNextMatch(func(have token.TokenType) bool {
		return want == have
	})
}

func (a *auditor) doesNextMatch(f func(token.TokenType) bool) bool {
	if !a.More() {
		return false
	}

	tk, e := a.Peek()
	if e != nil {
		e = err.AfterToken(a.prev, e, "Failed to read token")
		panic(e)
	}

	return f(tk.TokenType)
}

func (a *auditor) accept(want token.TokenType) bool {
	return a.acceptIf(func(have token.TokenType) bool {
		return want == have
	})
}

func (a *auditor) acceptIf(f func(token.TokenType) bool) bool {
	if !a.More() {
		return false
	}

	tk, e := a.Peek()
	if e != nil {
		e = err.AfterToken(a.prev, e, "Failed to read token")
		panic(e)
	}

	if !f(tk.TokenType) {
		return false
	}

	a.prev, e = a.Read()
	if e != nil {
		e = err.AtToken(tk, e, "Failed to read token")
		panic(e)
	}

	return true
}

func (a *auditor) expect(want token.TokenType) token.Token {
	return a.expectIf(func(have token.TokenType) bool {
		return want == have
	}, want.String())
}

func (a *auditor) expectIf(f func(token.TokenType) bool, exp any) token.Token {
	if !a.More() {
		e := err.AfterToken(a.prev, err.UnexpectedEOF, "Expected %q but got EOF", exp)
		panic(e)
	}

	tk, e := a.Read()
	if e != nil {
		e = err.AfterToken(a.prev, e, "Failed to read token")
		panic(e)
	}

	if !f(tk.TokenType) {
		e = err.AtToken(tk, err.UnexpectedToken, "Expected %q but got %q", exp, tk.TokenType)
		panic(e)
	}

	a.prev = tk
	return a.prev
}
