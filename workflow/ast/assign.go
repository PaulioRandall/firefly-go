package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type Assign struct {
	baseProc
	Token token.Token
	Left  Variable
	Right Expr
}

func (n Assign) Debug() string {
	return fmt.Sprintf("Assign %q\n\tLeft: %s\n\tRight: %s", n.Token.Value, n.Left.Debug(), n.Right.Debug())
}
