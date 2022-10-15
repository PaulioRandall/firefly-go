package ast

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// Node represents an AST, or put differently, a executable statement
//
// Nodes maybe nested and dependent upon others in an acyclic manner such that
// at runtime the dependees are executed first and their results used in the
// node's execution
type Node interface {
	node()

	// Where returns the start and end positions of
	// Where() (from, to pos.Pos)
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

// rangedNode is a node that can be mapped to a range of runes within a file
type rangedNode struct {
	from, to pos.Pos
}

func (n rangedNode) node() {}
func (n rangedNode) Where() (from, to pos.Pos) {
	return n.from, n.to
}

type Literal struct {
	rangedNode
	Operator token.Token
}

func (n Literal) node() {}
func (n Literal) stmt() {}
func (n Literal) proc() {}
func (n Literal) expr() {}

type Variable struct {
	Operator token.Token
}

func (n Variable) node() {}
func (n Variable) stmt() {}
func (n Variable) proc() {}
func (n Variable) expr() {}

type Assign struct {
	Left     []Variable
	Operator token.Token
	Right    []Expr
}

func (n Assign) node() {}
func (n Assign) stmt() {}
func (n Assign) proc() {}

type If struct {
	Keyword   token.Token
	Condition Expr
	Body      []Stmt
	End       token.Token
}

func (n If) node() {}
func (n If) stmt() {}

// TODO: Add to debug pkg
type When struct {
	Keyword token.Token
	Subject Expr
	Cases   []WhenCase
	End     token.Token
}

func (n When) node() {}
func (n When) stmt() {}

// TODO: Add to debug pkg
type WhenCase struct {
	Condition Expr
	Stmt      Stmt
}

func (n WhenCase) node() {}
