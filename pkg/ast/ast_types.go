package ast

type AstType int

const (
	AstTypeUndefined AstType = iota
	AstTypeNumber
	AstTypeOperation
)

var astTypeNames = map[AstType]string{
	AstTypeNumber:    "NUMBER",
	AstTypeOperation: "OPERATION",
}

func (at AstType) String() string {
	return astTypeNames[at]
}
