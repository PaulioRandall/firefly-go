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

type SeriesOfVar struct {
	Nodes []Variable
}

func (n SeriesOfVar) node() {}
func (n SeriesOfVar) stmt() {}
func (n SeriesOfVar) proc() {}

func (n SeriesOfVar) Where() (from, to pos.Pos) {
	lastIdx := len(n.Nodes) - 1
	from, _ = n.Nodes[0].Where()
	_, to = n.Nodes[lastIdx].Where()
	return from, to
}

type SeriesOfExpr struct {
	Nodes []Expr
}

func (n SeriesOfExpr) node() {}
func (n SeriesOfExpr) stmt() {}
func (n SeriesOfExpr) proc() {}

func (n SeriesOfExpr) Where() (from, to pos.Pos) {
	lastIdx := len(n.Nodes) - 1
	from, _ = n.Nodes[0].Where()
	_, to = n.Nodes[lastIdx].Where()
	return from, to
}
