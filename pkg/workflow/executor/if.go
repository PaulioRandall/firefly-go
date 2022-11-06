package executor

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func exeIf(state *exeState, n ast.If) {
	if exeExpr(state, n.Condition).(bool) {
		exeStmts(state, n.Body)
	}
}
