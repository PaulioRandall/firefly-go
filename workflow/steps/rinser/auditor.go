package rinser

import (
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type TokenReader interface {
	More() bool
	Read() token.Token
	Peek() token.Token
}

type auditor struct {
	TokenReader
	curr token.Token
}

func (a auditor) access() token.Token {
	return a.curr
}

func (a *auditor) accept(tt token.TokenType) bool {
	if a.More() && a.Peek().Type == tt {
		a.curr = a.Read()
		return true
	}
	return false
}

func (a *auditor) expect(tt token.TokenType) error {
	if !a.More() {
		return errAfter(a.curr, EOF, "Expected %q but got EOF", tt)
	}

	a.curr = a.Read()
	if tt != a.curr.Type {
		return errAfter(
			a.curr,
			UnexpectedToken,
			"Expected %q but got %q", tt, a.curr.Type)
	}

	return nil
}
