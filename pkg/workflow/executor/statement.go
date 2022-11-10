package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func exeStmts(mem *Memory, nodes []ast.Stmt) {
	for _, n := range nodes {
		exeNode(mem, n)
	}
}

func exeStmt(mem *Memory, n ast.Stmt) {
	switch v := n.(type) {
	case ast.Assign:
		exeAssign(mem, v)
	case ast.If:
		exeIf(mem, v)
	default:
		panic(ErrUnknownNode.Track("Unknown memment type"))
	}
}
