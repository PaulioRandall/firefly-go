package ast

import (
	"fmt"
)

// Node represents an AST, or put differently, a executable statement
//
// Nodes maybe nested and dependent upon others in an acyclic manner such that
// at runtime the dependees are executed first and their results used in the
// node's execution
type Node interface {
	node()
	Debug() string
	Print()
	Println()
}

// Stmt is a constraint for an executable statement
type Stmt interface {
	Node
	stmt()
}

// Proc (Procedure) is a constraint for a resolvable expression that can return
// any number of output values including none at all
type Proc interface {
	Stmt
	proc()
}

// Expr is specific Proc constraint that only and always returns a single value
//
// All sub nodes (recursive) of an Expr will also be an Expr
type Expr interface {
	Proc
	expr()
}

type baseNode struct{}

func (n baseNode) node()         {}
func (n baseNode) Debug() string { return "¯\\_(ツ)_/¯" }
func (n baseNode) Print()        { fmt.Print(n.Debug()) }
func (n baseNode) Println()      { fmt.Println(n.Debug()) }

type baseStmt struct{ baseNode }

func (n baseStmt) stmt() {}

type baseProc struct{ baseStmt }

func (n baseProc) proc() {}

type baseExpr struct{ baseProc }

func (n baseExpr) expr() {}
