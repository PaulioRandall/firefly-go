package ast

type ExprType int

const (
	ExprTypeUndefined ExprType = iota
	ExprTypeAdd
	ExprTypeSub
	ExprTypeMul
	ExprTypeDiv
)

var exprTypeNames = map[ExprType]string{
	ExprTypeAdd: "ADD",
	ExprTypeSub: "SUBTRACT",
	ExprTypeMul: "MULTIPLY",
	ExprTypeDiv: "DIVIDE",
}

func (et ExprType) String() string {
	return exprTypeNames[et]
}
