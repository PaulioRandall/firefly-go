package ast

import (
	"testing"
)

func Test_1_literal(t *testing.T) {
	_ = Expr(literal{})
}

func Test_2_variable(t *testing.T) {
	_ = Expr(variable{})
}

func Test_3_variable(t *testing.T) {
	_ = Proc(assign{})
}
