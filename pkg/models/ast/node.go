package ast

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// Node represents an Abstract Syntax Tree (AST) or executable statement
//
// Nodes maybe nested and dependent upon others in an acyclic manner such that
// at runtime any dependees are executed first and their results used in the
// node's execution.
type Node interface {
	// Where returns the start and end positions of the node within a source file
	Where() (from, to pos.Pos)
	node()
}

// Stmt represents a statement or constraint for an executable statement
type Stmt interface {
	Node
	stmt()
}

// Proc represents a statement that may be invoked as a procedure
//
// Invokation may return any number of output values including none at all.
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

// Literal represents one of the following literal tokens:
// - True
// - False
// - Number
// - String
type Literal struct {
	rangedNode
	Token token.Token
}

func (n Literal) stmt() {}
func (n Literal) proc() {}
func (n Literal) expr() {}

// Variable represents a variable value referenced using an identifier
type Variable struct {
	rangedNode
	Identifier token.Token
}

func (n Variable) stmt() {}
func (n Variable) proc() {}
func (n Variable) expr() {}

// Assign represents an assignment with left being the target variables and
// right being the statement that determines the new or updated variable values
type Assign struct {
	rangedNode
	Left     []Variable
	Operator token.Token
	Right    []Expr
}

func (n Assign) stmt() {}
func (n Assign) proc() {}

// If represents a conditional block of statements
type If struct {
	rangedNode
	Keyword   token.Token
	Condition Expr
	Body      []Stmt
	End       token.Token
}

func (n If) stmt() {}

// When represents a match block or branch with multiple options
type When struct {
	rangedNode
	Keyword token.Token
	Subject Expr
	Cases   []WhenCase
	End     token.Token
}

func (n When) stmt() {}

// WhenCase represents a matchable case within a When block
//
// Basically a sophisticated switch case without fallthrough
type WhenCase struct {
	rangedNode
	Condition Expr
	Statement Stmt
}
