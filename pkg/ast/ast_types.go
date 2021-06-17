package ast

type AstType int

const (
	AstTypeUndefined AstType = iota
	AstTypeNumber
	AstTypeAdd
	AstTypeSub
	AstTypeMul
	AstTypeDiv
)

var astTypeNames = map[AstType]string{
	AstTypeNumber: "NUMBER",
	AstTypeAdd:    "ADD",
	AstTypeSub:    "SUBTRACT",
	AstTypeMul:    "MULTIPLY",
	AstTypeDiv:    "DIVIDE",
}

func (at AstType) String() string {
	return astTypeNames[at]
}
