package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

type If struct {
	baseExpr
	keyword   token.Token
	condition Expr
	body      []Stmt
}

func MakeIf(keyword token.Token, condition Expr, body []Stmt) If {
	return If{
		keyword:   keyword,
		condition: condition,
		body:      body,
	}
}

func (n If) Debug() string {
	return fmt.Sprintf("If %q", n.keyword.Value)
}
