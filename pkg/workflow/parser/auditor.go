package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/container"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

var zero token.Token

type ReaderOfTokens = inout.Reader[token.Token]

type auditor struct {
	reader ReaderOfTokens
	buffer container.Queue[token.Token]
	prev   token.Token
}

func newAuditor(r ReaderOfTokens) *auditor {
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

	panic(err.NewPos(a.prev.To, "Failed to peek token from buffer"))
}

func (a *auditor) readNext() token.Token {
	a.loadBuffer()

	if tk, ok := a.buffer.Take(); ok {
		return tk
	}

	panic(err.NewPos(a.prev.To, "Failed to read token from buffer"))
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
		panic(err.WrapPos(e, a.prev.To, "Failed to read token from reader"))
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
	return a.acceptFunc(func(have token.TokenType) bool {
		return want == have
	})
}

func (a *auditor) acceptFunc(f func(token.TokenType) bool) bool {
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
	return a.expectFunc(want.String(), func(have token.TokenType) bool {
		return want == have
	})
}

func (a *auditor) expectFunc(exp any, f func(token.TokenType) bool) token.Token {
	if !a.more() {
		panic(err.WrapPosf(UnexpectedEOF, a.prev.To, "Expected %q but got EOF", exp))
	}

	tk := a.readNext()
	if !f(tk.TokenType) {
		panic(err.WrapPosf(UnexpectedToken, a.prev.To, "Expected %q but got %q", exp, tk.TokenType))
	}

	a.prev = tk
	return a.prev
}

func (a *auditor) expectWith(e error, want token.TokenType) token.Token {
	return a.expectFuncWith(e, func(have token.TokenType) bool {
		return want == have
	})
}

func (a *auditor) expectFuncWith(e error, f func(token.TokenType) bool) token.Token {
	if !a.more() {
		panic(err.WrapPosf(UnexpectedEOF, a.prev.To, "Failed to match token"))
	}

	tk := a.readNext()
	if !f(tk.TokenType) {
		panic(err.WrapPosf(e, a.prev.To, "Failed to match token"))
	}

	a.prev = tk
	return a.prev
}
