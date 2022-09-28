package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type variable struct {
	baseExpr
	tk token.Token
}

func MakeVariable(tk token.Token) variable {
	return variable{
		tk: tk,
	}
}

func (n variable) Debug() string {
	return fmt.Sprintf("Variable %q", n.tk.Value)
}
