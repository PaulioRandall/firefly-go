package ast

// Tree represents an AST
//
// Trees maybe nested and dependent upon others in an acyclic manner such that
// at runtime the dependees are executed first and their results used in the
// parents execution
type Tree interface {
	ASTType
	Is(ASTType) bool
}

// Stmt represents an executable statement
type Stmt interface {
	Tree
	stmt()
}

// Bloc represents an executable block of statements
type Bloc interface {
	Stmt
	bloc()
}

// Proc (Procedure) represents a resolvable expression that can return any
// number of values including zero
type Proc interface {
	Stmt
	proc()
}

// Expr represents a resolvable expression that always returns a single value
//
// All sub trees (recursive) of an Expr will also be an Expr or some further
// derived type
type Expr interface {
	Proc
	expr()
}

// MultiExpr represents an Expr that accepts any number of parameters
type MultiExpr interface {
	Expr
	multiExpr()
	Params() Expr
}

// BinaryExpr represents an Expr with exactly two parameters
type BinaryExpr interface {
	Expr
	binaryExpr()
	Left() Expr
	Right() Expr
}

// UnaryExpr represents an Expr with exactly one parameter
type UnaryExpr interface {
	Expr
	unaryExpr()
	Param() Expr
}
