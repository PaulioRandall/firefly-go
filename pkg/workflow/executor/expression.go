package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	//"github.com/PaulioRandall/firefly-go/pkg/utilities/debug"
)

func exeExpr(state *exeState, n ast.Expr) any {
	switch v := n.(type) {
	case ast.Variable:
		val, _ := state.getVariable(v.Name)
		return val
	case ast.Literal:
		return v.Value
	case ast.BinaryOperation:
		return exeBinaryOperation(state, v)
	default:
		panic(ErrUnknownNode.Track("Unknown expression type"))
	}
}
