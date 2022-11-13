package executor

import (
	"math"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/memory"
)

func exeBinaryOperation(mem *memory.Memory, n ast.BinaryOperation) any {

	left := exeExpr(mem, n.Left)
	right := exeExpr(mem, n.Right)

	switch n.Operator {
	case "+":
		return left.(float64) + right.(float64)
	case "-":
		return left.(float64) - right.(float64)
	case "*":
		return left.(float64) * right.(float64)
	case "/":
		return left.(float64) / right.(float64)
	case "%":
		return math.Mod(left.(float64), right.(float64))

	case "<":
		return left.(float64) < right.(float64)
	case ">":
		return left.(float64) > right.(float64)
	case "<=":
		return left.(float64) <= right.(float64)
	case ">=":
		return left.(float64) >= right.(float64)

	case "==":
		return left == right
	case "!=":
		return left != right

	case "&&":
		return left.(bool) && right.(bool)
	case "||":
		return left.(bool) || right.(bool)

	default:
		panic(ErrUnknownNode.Track("Unknown binary operator"))
	}
}
