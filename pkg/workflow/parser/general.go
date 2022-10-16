package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func notEndOfBlock(r PosReaderOfTokens) bool {
	return r.More() && !isNext(r, token.End)
}

func isNext(r PosReaderOfTokens, want token.TokenType) bool {
	if r.More() {
		return want == r.Peek().TokenType
	}
	return false
}

func doesNextMatch(r PosReaderOfTokens, f func(token.TokenType) bool) bool {
	if r.More() {
		return f(r.Peek().TokenType)
	}
	return false
}

func accept(r PosReaderOfTokens, want token.TokenType) bool {
	if !r.More() {
		return false
	}

	if want == r.Peek().TokenType {
		r.Read()
		return true
	}

	return false
}

func acceptFunc(r PosReaderOfTokens, f func(token.TokenType) bool) bool {
	if !r.More() {
		return false
	}

	if f(r.Peek().TokenType) {
		r.Read()
		return true
	}

	return false
}

func expect(r PosReaderOfTokens, want token.TokenType) token.Token {
	if !r.More() {
		panic(err.WrapPosf(UnexpectedEOF, r.Prev().To, "Expected %q but got EOF", want))
	}

	tk := r.Peek()
	if want == tk.TokenType {
		return r.Read()
	}

	panic(err.WrapPosf(UnexpectedToken, r.Prev().To, "Expected %q but got %q", want, tk.TokenType))
}

func expectFunc(r PosReaderOfTokens, want any, f func(token.TokenType) bool) token.Token {
	if !r.More() {
		panic(err.WrapPosf(UnexpectedEOF, r.Prev().To, "Expected %q but got EOF", want))
	}

	tk := r.Peek()
	if f(tk.TokenType) {
		return r.Read()
	}

	panic(err.WrapPosf(UnexpectedToken, r.Prev().To, "Expected %q but got %q", want, tk.TokenType))
}
