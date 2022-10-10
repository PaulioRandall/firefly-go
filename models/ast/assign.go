package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/models/token"
)

type Assign struct {
	baseProc
	Token token.Token
	Left  []Variable
	Right []Expr
}

func MakeAssign(tk token.Token, left []Variable, right []Expr) Assign {
	return Assign{
		Token: tk,
		Left:  left,
		Right: right,
	}
}

func (n Assign) Debug() string {
	return fmt.Sprintf("TODO: %q", n.Token.Value)
}
