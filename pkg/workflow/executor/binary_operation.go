package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func exeBinaryOperation(state *exeState, n ast.BinaryOperation) any {

	left := exeExpr(state, n.Left)
	right := exeExpr(state, n.Right)

	switch n.Operator {
	case "==":
		return left == right
	case "!=":
		return left != right
	case "<":
		return left.(float64) < right.(float64)
	case ">":
		return left.(float64) > right.(float64)
	case "<=":
		return left.(float64) <= right.(float64)
	case ">=":
		return left.(float64) >= right.(float64)
	case "&&":
		return left.(bool) && right.(bool)
	case "||":
		return left.(bool) || right.(bool)
	default:
		panic(ErrUnknownNode.Track("Unknown binary operator"))
	}
}
