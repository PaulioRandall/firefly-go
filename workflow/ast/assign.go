package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type assign struct {
	baseProc
	tk    token.Token
	left  variable
	right Expr
}

func MakeAssign(tk token.Token, left variable, right Expr) assign {
	return assign{
		tk:    tk,
		left:  left,
		right: right,
	}
}

func (n assign) Debug() string {
	return fmt.Sprintf("Literal %q", n.tk.Value)
}
