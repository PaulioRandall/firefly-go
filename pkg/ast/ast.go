package ast

type (
	Tree interface {
		Type() AstType
	}

	Number struct {
		Value int64
	}
)

func (t Number) Type() AstType { return AstTypeNumber }
