package ast

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

type If struct {
	baseExpr
	Keyword   token.Token
	Condition Expr
	Body      []Stmt
	End       token.Token
}

func MakeIf(
	keyword token.Token,
	condition Expr,
	body []Stmt,
	end token.Token,
) If {
	return If{
		Keyword:   keyword,
		Condition: condition,
		Body:      body,
		End:       end,
	}
}

func (n If) Debug() string {
	return fmt.Sprintf("If %q", n.Keyword.Value)
}
