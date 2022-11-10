package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	//"github.com/PaulioRandall/firefly-go/pkg/utilities/debug"
)

func exeExpr(mem *Memory, n ast.Expr) any {
	switch v := n.(type) {
	case ast.Variable:
		return mem.Variables[v.Name]
	case ast.Literal:
		return v.Value
	case ast.BinaryOperation:
		return exeBinaryOperation(mem, v)
	default:
		panic(ErrUnknownNode.Track("Unknown expression type"))
	}
}
