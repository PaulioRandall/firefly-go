package ast

// Tree represents an AST
//
// Trees maybe nested and dependent upon others in an acyclic manner
type Tree interface {
	Type() ASTType
	Is(ASTType) bool
}

// Stmt represents an executable statement
type Stmt interface {
	Tree
}

// Bloc represents an executable block of statements
type Bloc interface {
	Stmt
}

// Proc (Procedure) represents a resolvable expression that can return any
// number of values including zero
//
// All sub trees of a Proc will also be a Proc or some further derived type
type Proc interface {
	Stmt
}

// Expr represents a resolvable expression that always returns a single value
//
// All sub trees of an Expr will also be an Expr or some further derived type
type Expr interface {
	Stmt
}

// Proc returns t as a Proc or nil if not a valid subtype
func ToProc(t Tree) Proc {
	return nil
}

// Expr returns t as an Expr or nil if not a valid subtype
func ToExpr(t Tree) Expr {
	return nil
}
