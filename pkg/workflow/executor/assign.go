package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func exeAssign(state *exeState, n ast.Assign) {
	result := make([]any, len(n.Src))

	for i, v := range n.Src {
		result[i] = exeExpr(state, v)
	}

	for i, dst := range n.Dst {
		state.setVariable(dst.Name, result[i])
	}
}
