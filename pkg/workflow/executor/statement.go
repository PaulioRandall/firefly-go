package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func exeStmts(state *exeState, nodes []ast.Stmt) {
	for _, n := range nodes {
		exeNode(state, n)
	}
}

func exeStmt(state *exeState, n ast.Stmt) {
	switch v := n.(type) {
	case ast.Assign:
		exeAssign(state, v)
	case ast.If:
		exeIf(state, v)
	default:
		panic(ErrUnknownNode.Track("Unknown statement type"))
	}
}
