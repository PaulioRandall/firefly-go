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
	return nil
}
