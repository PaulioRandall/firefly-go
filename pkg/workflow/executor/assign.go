package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/memory"
)

func exeAssign(mem *memory.Memory, n ast.Assign) {
	result := make([]any, len(n.Src))

	for i, v := range n.Src {
		result[i] = exeExpr(mem, v)
	}

	for i, dst := range n.Dst {
		mem.Variables[dst.Name] = result[i]
	}
}
