package ast

type (
	Node interface{ Type() AST }

	Number struct{ Value int64 }

	InfixOperation struct {
		Left  Node
		Right Node
	}

	Add struct{ InfixOperation }
	Sub struct{ InfixOperation }
	Mul struct{ InfixOperation }
	Div struct{ InfixOperation }
)

func (t Number) Type() AST { return AstNumber }
func (t Add) Type() AST    { return AstAdd }
func (t Sub) Type() AST    { return AstSub }
func (t Mul) Type() AST    { return AstMul }
func (t Div) Type() AST    { return AstDiv }
