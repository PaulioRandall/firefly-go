package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/models/token"
)

type Variable struct {
	baseExpr
	Token token.Token
}

func MakeVariable(tk token.Token) Variable {
	return Variable{
		Token: tk,
	}
}

func (n Variable) Debug() string {
	return fmt.Sprintf("Variable %q", n.Token.Value)
}
