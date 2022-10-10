package parser

import (
	"github.com/PaulioRandall/firefly-go/workflow/container"
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var zero token.Token

type TokenReader = inout.Reader[token.Token]

type auditor struct {
	reader TokenReader
	buffer container.Queue[token.Token]
	prev   token.Token
}

func newAuditor(r TokenReader) *auditor {
	return &auditor{
		reader: r,
		buffer: &container.LinkedQueue[token.Token]{},
	}
}

func (a auditor) getPrev() token.Token {
	return a.prev
}

func (a *auditor) peekNext() token.Token {
	a.loadBuffer()

	if tk, ok := a.buffer.First(); ok {
		return tk
	}

	e := err.AfterToken(a.prev, nil, "Failed to peek token from buffer")
	panic(e)
}

func (a *auditor) readNext() token.Token {
	a.loadBuffer()

	if tk, ok := a.buffer.Take(); ok {
		return tk
	}

	e := err.AfterToken(a.prev, nil, "Failed to read token from buffer")
	panic(e)
}

func (a *auditor) putback(tk token.Token) {
	a.buffer.Return(tk)
}

func (a *auditor) loadBuffer() {
	if a.buffer.More() {
		return
	}

	tk, e := a.reader.Read()
	if e != nil {
		e = err.AfterToken(a.prev, e, "Failed to read token")
		panic(e)
	}

	a.buffer.Add(tk)
}

func (a *auditor) more() bool {
	return a.buffer.More() || a.reader.More()
}

func (a *auditor) isNext(want token.TokenType) bool {
	return a.doesNextMatch(func(have token.TokenType) bool {
		return want == have
	})
}

func (a *auditor) doesNextMatch(f func(token.TokenType) bool) bool {
	if !a.more() {
		return false
	}

	tk := a.peekNext()
	return f(tk.TokenType)
}

func (a *auditor) accept(want token.TokenType) bool {
	return a.acceptIf(func(have token.TokenType) bool {
		return want == have
	})
}

func (a *auditor) acceptIf(f func(token.TokenType) bool) bool {
	if !a.more() {
		return false
	}

	if !f(a.peekNext().TokenType) {
		return false
	}

	a.prev = a.readNext()
	return true
}

func (a *auditor) expect(want token.TokenType) token.Token {
	return a.expectIf(func(have token.TokenType) bool {
		return want == have
	}, want.String())
}

func (a *auditor) expectIf(f func(token.TokenType) bool, exp any) token.Token {
	if !a.more() {
		e := err.AfterToken(a.prev, err.UnexpectedEOF, "Expected %q but got EOF", exp)
		panic(e)
	}

	tk := a.readNext()
	if !f(tk.TokenType) {
		e := err.AtToken(tk, UnexpectedToken, "Expected %q but got %q", exp, tk.TokenType)
		panic(e)
	}

	a.prev = tk
	return a.prev
}
