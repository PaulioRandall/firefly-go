package ast

type NodeType int

const (
	Unknown NodeType = iota

	StmtBloc
	Assign

	FuncCall

	ExprCall
	Add
	Sub
	Mul
	Div
	Mod
)

var nameMap = map[NodeType]string{
	StmtBloc: "Statement Block",
	Assign:   "Assign",

	FuncCall: "Function Call",

	ExprCall: "Expression Call",
	Add:      "Add",
	Sub:      "Subtract",
	Mul:      "Multiply",
	Div:      "Divide",
	Mod:      "Remainder",
}

func (nt NodeType) String() string {
	return nameMap[nt]
}
