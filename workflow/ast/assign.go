package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type Assign struct {
	baseProc
	Token token.Token
	Left  []Variable
	Right []Expr
}

func (n Assign) Debug() string {
	return fmt.Sprintf("TODO: %q", n.Token.Value)
}
