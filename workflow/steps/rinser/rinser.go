// Package rinser removes whitespace and syntactic sugar from a list of tokens
package rinser

import (
	//"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type tokenReader interface {
	More() bool
	Read() token.Token
	Peek() token.Token
}
