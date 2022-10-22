package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

type auditor struct {
	r PosReaderOfTokens
}

func (a *auditor) More() bool {
	return a.r.More()
}

func (a *auditor) is(want token.TokenType) bool {
	if a.r.More() {
		return want == a.r.Peek().TokenType
	}
	return false
}

func (a *auditor) isNot(want token.TokenType) bool {
	if a.r.More() {
		return want != a.r.Peek().TokenType
	}
	return true
}

func (a *auditor) match(f func(token.TokenType) bool) bool {
	if a.r.More() {
		return f(a.r.Peek().TokenType)
	}
	return false
}

func (a *auditor) accept(want token.TokenType) bool {
	if !a.r.More() {
		return false
	}

	if want == a.r.Peek().TokenType {
		a.r.Read()
		return true
	}

	return false
}

func (a *auditor) acceptFunc(f func(token.TokenType) bool) bool {
	if !a.r.More() {
		return false
	}

	if f(a.r.Peek().TokenType) {
		a.r.Read()
		return true
	}

	return false
}

func (a *auditor) expect(want token.TokenType) token.Token {
	if !a.r.More() {
		panic(err.WrapPosf(UnexpectedEOF, a.r.Prev().To, "Expected %q but got EOF", want))
	}

	tk := a.r.Peek()
	if want == tk.TokenType {
		return a.r.Read()
	}

	panic(err.WrapPosf(UnexpectedToken, a.r.Prev().To, "Expected %q but got %q", want, tk.TokenType))
}

func (a *auditor) expectFunc(want any, f func(token.TokenType) bool) token.Token {
	if !a.r.More() {
		panic(err.WrapPosf(UnexpectedEOF, a.r.Prev().To, "Expected %q but got EOF", want))
	}

	tk := a.r.Peek()
	if f(tk.TokenType) {
		return a.r.Read()
	}

	panic(err.WrapPosf(UnexpectedToken, a.r.Prev().To, "Expected %q but got %q", want, tk.TokenType))
}
