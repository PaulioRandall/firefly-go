package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type literal struct {
	baseNode
	tk token.Token
}

func MakeLiteral(tk token.Token) literal {
	return literal{
		tk: tk,
	}
}

func (n literal) Debug() string {
	return fmt.Sprintf("Literal %q %#v", n.NodeType.String(), n.tk.Value)
}
