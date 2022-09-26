// Package rinser removes whitespace and syntactic sugar from a list of tokens
package rinser

import (
	//"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type TokenReader interface {
	More() bool
	Read() token.Token
	Peek() token.Token
}

type RinseFunc func() (tk token.Token, f RinseFunc, e error)

func New(r TokenReader) RinseFunc {

	// TODO: Steps:
	// 1. Read token
	// 2. If token is space
	// 3.   goto 1
	// 4. else return token

	return nil
}
