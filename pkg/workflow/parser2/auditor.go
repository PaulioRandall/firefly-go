package parser2

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

type auditor struct {
	r PosReaderOfTokens
}

func (a *auditor) More() bool {
	return a.r.More()
}

func (a *auditor) Peek() token.TokenType {
	return a.r.Peek().TokenType
}

func (a *auditor) Prev() token.Token {
	return a.r.Prev()
}

func (a *auditor) Read() token.Token {
	return a.r.Read()
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

func (a *auditor) isAny(wanted ...token.TokenType) bool {
	if !a.r.More() {
		return false
	}

	have := a.r.Peek().TokenType
	for _, want := range wanted {
		if want == have {
			return true
		}
	}

	return false
}

func (a *auditor) isNotAny(wanted ...token.TokenType) bool {
	return !a.isAny(wanted...)
}

func (a *auditor) match(f func(token.TokenType) bool) bool {
	if a.r.More() {
		return f(a.r.Peek().TokenType)
	}
	return false
}

func (a *auditor) notMatch(f func(token.TokenType) bool) bool {
	if a.r.More() {
		return f(a.r.Peek().TokenType)
	}
	return true
}

type Ranked interface{ Precedence() int }

func (a *auditor) hasPriorityOver(other Ranked) bool {
	return a.r.Peek().TokenType.Precedence() > other.Precedence()
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

func (a *auditor) acceptIf(f func(token.TokenType) bool) bool {
	if !a.r.More() {
		return false
	}

	if f(a.r.Peek().TokenType) {
		a.r.Read()
		return true
	}

	return false
}

func (a *auditor) acquire(want token.TokenType) (token.Token, bool) {
	if a.accept(want) {
		return a.Prev(), true
	}
	return token.Token{}, false
}

func (a *auditor) acquireIf(f func(token.TokenType) bool) (token.Token, bool) {
	if a.acceptIf(f) {
		return a.Prev(), true
	}
	return token.Token{}, false
}

func (a *auditor) expect(want token.TokenType) token.Token {
	if !a.r.More() {
		panic(a.unexpectedEOF(want))
	}

	tk := a.r.Peek()
	if want == tk.TokenType {
		return a.r.Read()
	}

	panic(a.unexpectedToken(want, tk.TokenType))
}

func (a *auditor) expectFor(want any, f func(token.TokenType) bool) token.Token {
	if !a.r.More() {
		panic(a.unexpectedEOF(want))
	}

	tk := a.r.Peek()
	if f(tk.TokenType) {
		return a.r.Read()
	}

	panic(a.unexpectedToken(want, tk.TokenType))
}

// ******* NEW ********

func (a *auditor) expect_new(want token.TokenType) (token.Token, error) {
	var zero token.Token

	if !a.r.More() {
		return zero, a.unexpectedEOF(want)
	}

	tk := a.r.Peek()
	if want == tk.TokenType {
		return a.r.Read(), nil
	}

	return zero, a.unexpectedToken(want, tk.TokenType)
}

func (a *auditor) unexpectedToken(expected, got any) error {
	return ErrUnexpectedToken.Trackf("Expected token %q but got %q", expected, got)
}

func (a *auditor) unexpectedEOF(expected any) error {
	return ErrUnexpectedEOF.Trackf("Expected token %q but got EOF", expected)
}
