package ast

import (
	"testing"
)

func Test_enforceTypes(t *testing.T) {
	_ = Expr(Literal{})
	_ = Expr(Variable{})
	_ = Proc(Assign{})
	_ = Stmt(If{})
}
