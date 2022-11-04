package executor

import (
	"github.com/PaulioRandall/firefly-go/pkg/workflow/executor/ast"
)

func exeNode(state *exeState, n ast.Node) {
	switch v := n.(type) {
	case ast.Assign:
		exeAssign(state, v)
	default:
		panic("TODO: Unknown node")
	}
}

func exeExpr(state *exeState, n ast.Expr) any {
	switch v := n.(type) {
	case ast.Literal:
		return v.Value
	default:
		panic("TODO: Unknown expr")
	}
}

func exeAssign(state *exeState, n ast.Assign) {
	result := make([]any, len(n.Src))

	for i, v := range n.Src {
		result[i] = exeExpr(state, v)
	}

	for i, dst := range n.Dst {
		state.setVariable(dst.Name, result[i])
	}
}
