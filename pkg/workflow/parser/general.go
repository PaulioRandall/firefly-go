package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func notEndOfBlock(a tokenAuditor) bool {
	return a.More() && !isNext(a, token.End)
}

func isNext(a tokenAuditor, want token.TokenType) bool {
	if a.More() {
		return want == a.Peek().TokenType
	}
	return false
}

func doesNextMatch(a tokenAuditor, f func(token.TokenType) bool) bool {
	if a.More() {
		return f(a.Peek().TokenType)
	}
	return false
}

func accept(a tokenAuditor, want token.TokenType) bool {
	if !a.More() {
		return false
	}

	if want == a.Peek().TokenType {
		a.Read()
		return true
	}

	return false
}

func acceptFunc(a tokenAuditor, f func(token.TokenType) bool) bool {
	if !a.More() {
		return false
	}

	if f(a.Peek().TokenType) {
		a.Read()
		return true
	}

	return false
}

func expect(a tokenAuditor, want token.TokenType) token.Token {
	if !a.More() {
		panic(err.WrapPosf(UnexpectedEOF, a.Prev().To, "Expected %q but got EOF", want))
	}

	tk := a.Peek()
	if want == tk.TokenType {
		return a.Read()
	}

	panic(err.WrapPosf(UnexpectedToken, a.Prev().To, "Expected %q but got %q", want, tk.TokenType))
}

func expectFunc(a tokenAuditor, want any, f func(token.TokenType) bool) token.Token {
	if !a.More() {
		panic(err.WrapPosf(UnexpectedEOF, a.Prev().To, "Expected %q but got EOF", want))
	}

	tk := a.Peek()
	if f(tk.TokenType) {
		return a.Read()
	}

	panic(err.WrapPosf(UnexpectedToken, a.Prev().To, "Expected %q but got %q", want, tk.TokenType))
}
