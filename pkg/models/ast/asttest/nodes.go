package asttest

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func Literal(tk token.Token) ast.Literal {
	return ast.Literal{
		Token: tk,
	}
}

func Variable(id token.Token) ast.Variable {
	return ast.Variable{
		Identifier: id,
	}
}

func Assign(left []ast.Variable, op token.Token, right ast.Stmt) ast.Assign {
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

func WhenCase(
	condition ast.Expr,
	stmt ast.Stmt,
) ast.WhenCase {
	return ast.WhenCase{
		Condition: condition,
		Statement: stmt,
	}
}

func Is(
	keyword token.Token,
	expr ast.Expr,
) ast.Is {
	return ast.Is{
		Keyword: keyword,
		Expr:    expr,
	}
}

func ExprSet(exprs ...ast.Expr) ast.ExprSet {
	return ast.ExprSet{
		Exprs: exprs,
	}
}

func BinaryOperation(
	left ast.Expr,
	op token.Token,
	right ast.Expr,
) ast.BinaryOperation {
	return ast.BinaryOperation{
		Left:     left,
		Operator: op,
		Right:    right,
	}
}
