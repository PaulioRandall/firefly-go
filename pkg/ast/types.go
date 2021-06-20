package ast

type AST int

const (
	AstUndefined AST = iota
	AstEmpty
	AstNumber
	AstAdd
	AstSub
	AstMul
	AstDiv
)

var astNames = map[AST]string{
	AstEmpty:  "EMPTY",
	AstNumber: "NUMBER",
	AstAdd:    "ADD",
	AstSub:    "SUBTRACT",
	AstMul:    "MULTIPLY",
	AstDiv:    "DIVIDE",
}

func (a AST) String() string {
	return astNames[a]
}
