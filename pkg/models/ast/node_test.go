package ast

import (
	"testing"
)

func Test_enforceTypes(t *testing.T) {
	_ = Stmt(If{})
	_ = Stmt(When{})

	_ = Proc(Assign{})

	_ = Expr(Literal{})
	_ = Expr(Variable{})
}
