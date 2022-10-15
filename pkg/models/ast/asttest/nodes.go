package asttest

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func Literal(op token.Token) ast.Literal {
	return ast.Literal{
		Operator: op,
	}
}

func Variable(op token.Token) ast.Variable {
	return ast.Variable{
		Operator: op,
	}
}

func Assign(left []ast.Variable, op token.Token, right []ast.Expr) ast.Assign {
	return ast.Assign{
		Left:     left,
		Operator: op,
		Right:    right,
	}
}

func If(
	keyword token.Token,
	condition ast.Expr,
	body []ast.Stmt,
	end token.Token,
) ast.If {
	return ast.If{
		Keyword:   keyword,
		Condition: condition,
		Body:      body,
		End:       end,
	}
}

func When(
	keyword token.Token,
	subject ast.Expr,
	cases []ast.WhenCase,
	end token.Token,
) ast.When {
	return ast.When{
		Keyword: keyword,
		Subject: subject,
		Cases:   cases,
		End:     end,
	}
}
