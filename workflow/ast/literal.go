package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type Literal struct {
	baseExpr
	Token token.Token
}

func (n Literal) Debug() string {
	return fmt.Sprintf("Literal %q", n.Token.Value)
}
