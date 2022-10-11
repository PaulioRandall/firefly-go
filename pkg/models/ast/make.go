package ast

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func MakeLiteral(op token.Token) Literal {
	return Literal{
		Operator: op,
	}
}

func MakeVariable(op token.Token) Variable {
	return Variable{
		Operator: op,
	}
}

func MakeAssign(left []Variable, op token.Token, right []Expr) Assign {
	return Assign{
		Left:     left,
		Operator: op,
		Right:    right,
	}
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
