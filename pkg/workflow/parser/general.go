package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
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

func accept(a *auditor.Auditor, want token.TokenType) bool {
	if !a.More() {
		return false
	}

	if want == a.Peek().TokenType {
		a.Read()
		return true
	}

	return false
}

func acceptFunc(a *auditor.Auditor, f func(token.TokenType) bool) bool {
	if !a.More() {
		return false
	}

	if f(a.Peek().TokenType) {
		a.Read()
		return true
	}

	return false
}

func expect(a *auditor.Auditor, want token.TokenType) token.Token {
	if !a.More() {
		panic(err.WrapPosf(UnexpectedEOF, a.Prev().To, "Expected %q but got EOF", want))
	}

	tk := a.Peek()
	if want == tk.TokenType {
		return a.Read()
	}

	panic(err.WrapPosf(UnexpectedToken, a.Prev().To, "Expected %q but got %q", want, tk.TokenType))
}

func expectFunc(a *auditor.Auditor, want any, f func(token.TokenType) bool) token.Token {
	if !a.More() {
		panic(err.WrapPosf(UnexpectedEOF, a.Prev().To, "Expected %q but got EOF", want))
	}

	tk := a.Peek()
	if f(tk.TokenType) {
		return a.Read()
	}

	panic(err.WrapPosf(UnexpectedToken, a.Prev().To, "Expected %q but got %q", want, tk.TokenType))
}
