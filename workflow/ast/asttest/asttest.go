package asttest

import (
	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func Assign(tk token.Token, left []ast.Variable, right []ast.Expr) ast.Assign {
	return ast.Assign{
		Token: tk,
		Left:  left,
		Right: right,
	}
}

func Variable(tk token.Token) ast.Variable {
	return ast.Variable{
		Token: tk,
	}
}

func Literal(tk token.Token) ast.Literal {
	return ast.Literal{
		Token: tk,
	}
}
