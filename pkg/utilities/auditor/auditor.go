package auditor

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/container"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

var (
	zero token.Token

	UnexpectedEOF   = err.New("Unexpected end of file")
	UnexpectedToken = err.New("Unexpected token")
)

type ReaderOfTokens = inout.Reader[token.Token]

type Auditor struct {
	reader ReaderOfTokens
	buffer container.Queue[token.Token]
	prev   token.Token
}

func NewAuditor(r ReaderOfTokens) *Auditor {
	return &Auditor{
		reader: r,
		buffer: &container.LinkedQueue[token.Token]{},
	}
}

func (a Auditor) Prev() token.Token {
	return a.prev
}

func (a *Auditor) More() bool {
	return a.buffer.More() || a.reader.More()
}

func (a *Auditor) Peek() token.Token {
	a.loadBuffer()

	if tk, ok := a.buffer.First(); ok {
		return tk
	}

	panic(err.NewPos(a.prev.To, "Failed to peek token from buffer"))
}

func (a *Auditor) Read() token.Token {
	a.loadBuffer()

	if tk, ok := a.buffer.Take(); ok {
		return tk
	}

	panic(err.NewPos(a.prev.To, "Failed to read token from buffer"))
}

func (a *Auditor) Putback(tk token.Token) {
	a.buffer.Return(tk)
}

func (a *Auditor) loadBuffer() {
	if a.buffer.More() {
		return
	}

	tk, e := a.reader.Read()
	if e != nil {
		panic(err.WrapPos(e, a.prev.To, "Failed to read token from reader"))
	}

	a.buffer.Add(tk)
}

func (a *Auditor) IsNext(want token.TokenType) bool {
	return a.DoesNextMatch(func(have token.TokenType) bool {
		return want == have
	})
}

func (a *Auditor) DoesNextMatch(f func(token.TokenType) bool) bool {
	if !a.More() {
		return false
	}

	tk := a.Peek()
	return f(tk.TokenType)
}

func (a *Auditor) Accept(want token.TokenType) bool {
	return a.AcceptFunc(func(have token.TokenType) bool {
		return want == have
	})
}

func (a *Auditor) AcceptFunc(f func(token.TokenType) bool) bool {
	if !a.More() {
		return false
	}

	if !f(a.Peek().TokenType) {
		return false
	}

	a.prev = a.Read()
	return true
}

func (a *Auditor) Expect(want token.TokenType) token.Token {
	return a.ExpectFunc(want.String(), func(have token.TokenType) bool {
		return want == have
	})
}

func (a *Auditor) ExpectFunc(exp any, f func(token.TokenType) bool) token.Token {
	if !a.More() {
		panic(err.WrapPosf(UnexpectedEOF, a.prev.To, "Expected %q but got EOF", exp))
	}

	tk := a.Read()
	if !f(tk.TokenType) {
		panic(err.WrapPosf(UnexpectedToken, a.prev.To, "Expected %q but got %q", exp, tk.TokenType))
	}

	a.prev = tk
	return a.prev
}
