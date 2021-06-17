package ast

type (
	Node interface {
		Type() AST
	}

	Number struct {
		Value int64
	}
)

func (t Number) Type() AST { return AstNumber }
