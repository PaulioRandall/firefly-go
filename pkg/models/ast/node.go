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
	Precedence() int
}

// Literal represents one of the following literal tokens:
// - True
// - False
// - Number
// - String
type Literal struct {
	Token token.Token
}

func (n Literal) node() {}
func (n Literal) stmt() {}
func (n Literal) proc() {}
func (n Literal) expr() {}
func (n Literal) Precedence() int {
	return n.Token.Precedence()
}
func (n Literal) Where() (from, to pos.Pos) {
	return n.Token.Where()
}

// Variable represents a variable value referenced using an identifier
type Variable struct {
	Identifier token.Token
}

func (n Variable) node() {}
func (n Variable) stmt() {}
func (n Variable) proc() {}
func (n Variable) expr() {}
func (n Variable) Precedence() int {
	return n.Identifier.Precedence()
}
func (n Variable) Where() (from, to pos.Pos) {
	return n.Identifier.Where()
}

// BinaryOperation represents an expression with two operators
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

// Assign represents an assignment with left being the target variables and
// right being the statement that determines the new or updated variable values
type Assign struct {
	Left     []Variable
	Operator token.Token
	Right    Stmt
}

func (n Assign) node() {}
func (n Assign) stmt() {}
func (n Assign) Where() (from, to pos.Pos) {
	from, _ = n.Left[0].Where()
	_, to = n.Right.Where()
	return from, to
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

// If represents a conditional block of statements
type If struct {
	Keyword   token.Token
	Condition Expr
	Body      []Stmt
	End       token.Token
}

func (n If) node() {}
func (n If) stmt() {}
func (n If) Where() (from, to pos.Pos) {
	from, _ = n.Keyword.Where()
	_, to = n.End.Where()
	return from, to
}

// When represents a match block or branch with multiple options
type When struct {
	Keyword token.Token
	Subject Expr
	Cases   []WhenCase
	End     token.Token
}

func (n When) node() {}
func (n When) stmt() {}
func (n When) Where() (from, to pos.Pos) {
	from, _ = n.Keyword.Where()
	_, to = n.End.Where()
	return from, to
}

// WhenCase represents a matchable case within a When block
//
// Basically a sophisticated switch case without fallthrough
type WhenCase struct {
	Condition Expr
	Statement Stmt
}

func (n WhenCase) node() {}
func (n WhenCase) Where() (from, to pos.Pos) {
	from, _ = n.Condition.Where()
	_, to = n.Statement.Where()
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
