package executor

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

var (
	ErrUnknownNode = err.Trackable("Unknown Node")
)

func exeNodes(state *exeState, nodes []ast.Node) {
	for _, n := range nodes {
		exeNode(state, n)
	}
}

func exeNode(state *exeState, n ast.Node) {
	switch v := n.(type) {
	case ast.Stmt:
		exeStmt(state, v)
	default:
		panic(ErrUnknownNode.Track("Unknown node type"))
	}
}
