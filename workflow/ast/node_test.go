package ast

import (
	"testing"
)

func asStmt[T Stmt](v T) T { return v }
func asProc[T Proc](v T) T { return v }
func asExpr[T Expr](v T) T { return v }

func Test_1_literal(t *testing.T) {
	_ = Node(literal{})
	_ = asExpr(literal{})
}
