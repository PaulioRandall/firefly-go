package parser

import (
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type TokenReader = inout.Reader[token.Token]

type auditor struct {
	TokenReader
	curr token.Token
}

func (a auditor) access() token.Token {
	return a.curr
}

func (a *auditor) accept(tt token.TokenType) bool {
	if !a.More() {
		return false
	}

	tk, e := a.Peek()
	if e != nil {
		e = err.AfterToken(a.curr, e, "Failed to read token")
		panic(e)
	}

	if tk.TokenType != tt {
		return false
	}

	a.curr, e = a.Read()
	if e != nil {
		e = err.AtToken(tk, e, "Failed to read token")
		panic(e)
	}

	return true
}

func (a *auditor) expect(tt token.TokenType) token.Token {
	if !a.More() {
		e := err.AfterToken(a.curr, err.UnexpectedEOF, "Expected %q but got EOF", tt)
		panic(e)
	}

	tk, e := a.Read()
	if e != nil {
		e = err.AfterToken(a.curr, e, "Failed to read token")
		panic(e)
	}

	if tk.TokenType != tt {
		e = err.AtToken(tk, err.UnexpectedToken, "Expected %q but got %q", tt, tk.TokenType)
		panic(e)
	}

	a.curr = tk
	return a.curr
}
