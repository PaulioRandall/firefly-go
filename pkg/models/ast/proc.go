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

type SeriesOfExpr struct {
	Exprs []Expr
}

func (n SeriesOfExpr) node() {}
func (n SeriesOfExpr) stmt() {}
func (n SeriesOfExpr) proc() {}

func (n SeriesOfExpr) Where() (from, to pos.Pos) {
	lastIdx := len(n.Exprs) - 1
	from, _ = n.Exprs[0].Where()
	_, to = n.Exprs[lastIdx].Where()
	return from, to
}
