package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
)

func notEndOfBlock(a *auditor.Auditor) bool {
	return a.More() && !isNext(a, token.End)
}

func isNext(a *auditor.Auditor, want token.TokenType) bool {
	if a.More() {
		return want == a.Peek().TokenType
	}
	return false
}

func doesNextMatch(a *auditor.Auditor, f func(token.TokenType) bool) bool {
	if a.More() {
		return f(a.Peek().TokenType)
	}
	return false
}

/*
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
*/
