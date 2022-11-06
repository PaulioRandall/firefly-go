package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func exeExpr(state *exeState, n ast.Expr) any {
	switch v := n.(type) {
	case ast.Literal:
		return v.Value
	case ast.BinaryOperation:
		return exeBinaryOperation(state, v)
	default:
		panic(ErrUnknownNode.Track("Unknown expression type"))
	}
}
