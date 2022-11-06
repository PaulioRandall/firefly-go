package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func litNumber(v float64) ast.Literal {
	return ast.Literal{Value: v}
}

func exeBinaryOperation(state *exeState, n ast.BinaryOperation) any {

	left := exeExpr(state, n.Left)
	right := exeExpr(state, n.Right)

	switch n.Operator {
	case "==":
		return left == right
	case "!=":
		return left != right
	}

	return nil
}
