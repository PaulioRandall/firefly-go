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

func (a *auditor) Prev() token.Token {
	return a.r.Prev()
}

func (a *auditor) Putback(tk token.Token) {
	a.r.Putback(tk)
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
		panic(a.unexpectedEOF(want))
	}

	tk := a.r.Peek()
	if want == tk.TokenType {
		return a.r.Read()
	}

	panic(a.unexpected(want, tk.TokenType))
}

func (a *auditor) expectFunc(want any, f func(token.TokenType) bool) token.Token {
	if !a.r.More() {
		panic(a.unexpectedEOF(want))
	}

	tk := a.r.Peek()
	if f(tk.TokenType) {
		return a.r.Read()
	}

	panic(a.unexpected(want, tk.TokenType))
}

func (a *auditor) unexpected(expected, got any) error {
	return a.wrapErr(UnexpectedToken, "Expected %q but got %q", expected, got)
}

func (a *auditor) unexpectedEOF(expected any) error {
	return a.wrapErr(UnexpectedEOF, "Expected %q but got EOF", expected)
}

func (a *auditor) wrapErr(cause error, msg string, args ...any) error {
	return err.WrapPosf(cause, a.r.Prev().To, msg, args...)
}
