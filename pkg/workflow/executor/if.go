package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/memory"
)

func exeIf(mem *memory.Memory, n ast.If) {
	if exeExpr(mem, n.Condition).(bool) {
		exeStmts(mem, n.Body)
	}
}
