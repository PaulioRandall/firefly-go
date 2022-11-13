package executor

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/memory"
)

var (
	ErrUnknownNode = err.Trackable("Unknown Node")
)

func exeNodes(mem *memory.Memory, nodes []ast.Node) {
	for _, n := range nodes {
		exeNode(mem, n)
	}
}

func exeNode(mem *memory.Memory, n ast.Node) {
	switch v := n.(type) {
	case ast.Stmt:
		exeStmt(mem, v)
	default:
		panic(ErrUnknownNode.Track("Unknown node type"))
	}
}
