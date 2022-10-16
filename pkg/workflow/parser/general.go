package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func notEndOfBlock(r BufReaderOfTokens) bool {
	return r.More() && !isNext(r, token.End)
}

func isNext(r BufReaderOfTokens, want token.TokenType) bool {
	if r.More() {
		return want == r.Peek().TokenType
	}
	return false
}

func doesNextMatch(r BufReaderOfTokens, f func(token.TokenType) bool) bool {
	if r.More() {
		return f(r.Peek().TokenType)
	}
	return false
}

func accept(r BufReaderOfTokens, want token.TokenType) bool {
	if !r.More() {
		return false
	}

	if want == r.Peek().TokenType {
		r.Read()
		return true
	}

	return false
}

func acceptFunc(r BufReaderOfTokens, f func(token.TokenType) bool) bool {
	if !r.More() {
		return false
	}

	if f(r.Peek().TokenType) {
		r.Read()
		return true
	}

	return false
}

func expect(r BufReaderOfTokens, want token.TokenType) token.Token {
	if !r.More() {
		panic(err.WrapPosf(UnexpectedEOF, r.Prev().To, "Expected %q but got EOF", want))
	}

	tk := r.Peek()
	if want == tk.TokenType {
		return r.Read()
	}

	panic(err.WrapPosf(UnexpectedToken, r.Prev().To, "Expected %q but got %q", want, tk.TokenType))
}

func expectFunc(r BufReaderOfTokens, want any, f func(token.TokenType) bool) token.Token {
	if !r.More() {
		panic(err.WrapPosf(UnexpectedEOF, r.Prev().To, "Expected %q but got EOF", want))
	}

	tk := r.Peek()
	if f(tk.TokenType) {
		return r.Read()
	}

	panic(err.WrapPosf(UnexpectedToken, r.Prev().To, "Expected %q but got %q", want, tk.TokenType))
}
