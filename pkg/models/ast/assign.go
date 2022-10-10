package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

type Assign struct {
	baseProc
	Left     []Variable
	Operator token.Token
	Right    []Expr
}

func MakeAssign(left []Variable, op token.Token, right []Expr) Assign {
	return Assign{
		Left:     left,
		Operator: op,
		Right:    right,
	}
}

func (n Assign) Debug() string {
	return fmt.Sprintf("TODO: %q", n.Operator.Value)
}
