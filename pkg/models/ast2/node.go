package ast2

// Node represents an Abstract Syntax Tree (AST) or executable statement.
//
// Nodes maybe nested and dependent upon others in an acyclic manner such that
// at runtime any dependees are executed first and their results used in the
// node's execution.
type Node interface {
	node()
}

type Stmt interface {
	Node
	stmt()
}

type Returns interface {
	Node
	returns()
}

type Expr interface {
	Returns
	expr()
}

type Assign struct {
	Dst []Variable
	Src []Expr // TODO: should be Returns
}

func (n Assign) node() {}
func (n Assign) stmt() {}

type If struct {
	Condition Expr
	Body      []Stmt
}

func (n If) node() {}
func (n If) stmt() {}

type Variable struct {
	Name string
}

func (n Variable) node()    {}
func (n Variable) expr()    {}
func (n Variable) returns() {}

type Literal struct {
	Value any
}

func (n Literal) node()    {}
func (n Literal) expr()    {}
func (n Literal) returns() {}

type BinaryOperation struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (n BinaryOperation) node()    {}
func (n BinaryOperation) expr()    {}
func (n BinaryOperation) returns() {}

type SpellCall struct {
	Name   string
	Params []Expr
}

func (n SpellCall) node()    {}
func (n SpellCall) stmt()    {}
func (n SpellCall) returns() {}
