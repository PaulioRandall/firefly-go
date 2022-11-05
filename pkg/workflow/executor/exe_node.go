package executor

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"

	"github.com/PaulioRandall/firefly-go/pkg/workflow/executor/ast"
)

var (
	ErrUnknownNode     = err.Trackable("Unknown Node")
	ErrUnknownExprNode = err.Trackable("Unknown Expr Node")
)

func exeNode(state *exeState, n ast.Node) {
	switch v := n.(type) {
	case ast.Assign:
		exeAssign(state, v)
	default:
		panic(unknownNode(nil))
	}
}

func exeExpr(state *exeState, n ast.Expr) any {
	switch v := n.(type) {
	case ast.Literal:
		return v.Value
	default:
		panic(unknownExprNode())
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

func unknownNode(e error) error {
	return ErrUnknownNode.Wrap(e, "Node type does not match any known executable type")
}

func unknownExprNode() error {
	e := ErrUnknownExprNode.Track("Expr type does not match any known executable type")
	return unknownNode(e)
}
