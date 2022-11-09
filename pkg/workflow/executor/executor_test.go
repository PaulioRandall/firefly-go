package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func mockVariable(v string) ast.Variable {
	return ast.Variable{Name: v}
}

func mockBool(v bool) ast.Literal {
	return ast.Literal{Value: v}
}

func mockNumber(v float64) ast.Literal {
	return ast.Literal{Value: v}
}

func mockString(v string) ast.Literal {
	return ast.Literal{Value: v}
}

func mockVariables(names ...string) []ast.Variable {
	n := make([]ast.Variable, len(names))

	for i, v := range names {
		n[i] = ast.Variable{
			Name: v,
		}
	}

	return n
}

func mockLiterals(values ...any) []ast.Expr {
	n := make([]ast.Expr, len(values))

	for i, v := range values {
		n[i] = ast.Literal{
			Value: v,
		}
	}

	return n
}
