package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type Variable struct {
	baseExpr
	Token token.Token
}

func (n Variable) Debug() string {
	return fmt.Sprintf("Variable %q", n.Token.Value)
}
