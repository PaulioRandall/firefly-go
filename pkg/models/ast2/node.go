package ast2

// Node represents an Abstract Syntax Tree (AST) or executable statement.
//
// Nodes maybe nested and dependent upon others in an acyclic manner such that
// at runtime any dependees are executed first and their results used in the
// node's execution.
type Node interface {
	node()
}

type Expr interface {
	expr()
}

type Assign struct {
	Dst []Variable
	Src []Expr
}

func (n Assign) node() {}

type Variable struct {
	Name string
}

func (n Variable) node() {}

type Literal struct {
	Value any
}

func (n Literal) node() {}
func (n Literal) expr() {}

type If struct {
	Condition Expr
	Body      []Node
}

func (n If) node() {}
