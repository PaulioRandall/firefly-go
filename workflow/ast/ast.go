package ast

// Node represents an AST
//
// Nodes maybe nested and dependent upon others in an acyclic manner such that
// at runtime the dependees are executed first and their results used in the
// parents execution
type Node interface {
	Type() ASTType
	Is(ASTType) bool
}

// Stmt represents an executable statement
type Stmt interface {
	Node
}

// Bloc represents an executable block of statements
type Bloc interface {
	Node
}

// Proc (Procedure) represents a resolvable expression that can return any
// number of values including zero
type Proc interface {
	Node
}

// Expr represents a resolvable expression that always returns a single value
//
// All sub nodes (recursive) of an Expr will also be an Expr or some further
// derived type
type Expr interface {
	Node
}

// MultiExpr represents an Expr that accepts any number of parameters
type MultiExpr interface {
	Expr
	Params() Expr
}

// BinaryExpr represents an Expr with exactly two parameters
type BinaryExpr interface {
	Expr
	Left() Expr
	Right() Expr
}

// UnaryExpr represents an Expr with exactly one parameter
type UnaryExpr interface {
	Expr
	Param() Expr
}
