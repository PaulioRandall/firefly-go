package ast

import (
	"testing"
)

func Test_1_literal(t *testing.T) {
	_ = Expr(Literal{})
}

func Test_2_variable(t *testing.T) {
	_ = Expr(Variable{})
}

func Test_3_assign(t *testing.T) {
	_ = Proc(Assign{})
}
