package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/memory"
)

func exeStmts(mem *memory.Memory, nodes []ast.Stmt) {
	for _, n := range nodes {
		exeNode(mem, n)
	}
}

func exeStmt(mem *memory.Memory, n ast.Stmt) {
	switch v := n.(type) {
	case ast.Assign:
		exeAssign(mem, v)
	case ast.If:
		exeIf(mem, v)
	case ast.SpellCall:
		invokeSpell(mem, v)
	default:
		panic(ErrUnknownNode.Track("Unknown memment type"))
	}
}
