package ast

type ASTType int

const (
	Unknown ASTType = iota

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
