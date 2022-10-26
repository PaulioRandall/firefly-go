package ast

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
)

// Proc represents a statement that may be invoked as a procedure
//
// Invokation may return any number of output values including none at all.
type Proc interface {
	Stmt
	proc()
}

type ExprSet struct {
	Exprs []Expr
}

func (n ExprSet) node() {}
func (n ExprSet) stmt() {}
func (n ExprSet) proc() {}

func (n ExprSet) Where() (from, to pos.Pos) {
	lastIdx := len(n.Exprs) - 1
	from, _ = n.Exprs[0].Where()
	_, to = n.Exprs[lastIdx].Where()
	return from, to
}
