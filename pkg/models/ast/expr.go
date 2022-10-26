package ast

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// Expr is specific Proc constraint that only and always returns a single value
//
// All sub nodes (recursive) of an Expr will also be an Expr
type Expr interface {
	Proc
	expr()
	Precedence() int
}

// BinaryOperation represents an expression with two operators.
type BinaryOperation struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (n BinaryOperation) node() {}
func (n BinaryOperation) stmt() {}
func (n BinaryOperation) proc() {}
func (n BinaryOperation) expr() {}

func (n BinaryOperation) Precedence() int {
	return n.Operator.Precedence()
}

func (n BinaryOperation) Where() (from, to pos.Pos) {
	from, _ = n.Left.Where()
	_, to = n.Right.Where()
	return from, to
}

// Is represents a when case which tests for equality between the When subject
// and the result of an expression.
type Is struct {
	Keyword token.Token
	Expr    Expr
}

func (n Is) node() {}
func (n Is) stmt() {}
func (n Is) proc() {}
func (n Is) expr() {}

func (n Is) Precedence() int {
	return 0
}

func (n Is) Where() (from, to pos.Pos) {
	from, _ = n.Keyword.Where()
	_, to = n.Expr.Where()
	return from, to
}
