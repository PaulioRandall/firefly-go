package ast

type (
	Ast  interface{ Type() AstType }
	Expr interface{ ExprType() ExprType }
)

type astExpr struct{ AstType AstType }

func (e astExpr) Type() AstType { return e.AstType }

type Number struct {
	astExpr
	Value int64
}
