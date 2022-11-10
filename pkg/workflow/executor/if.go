package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func exeIf(mem *Memory, n ast.If) {
	if exeExpr(mem, n.Condition).(bool) {
		exeStmts(mem, n.Body)
	}
}
