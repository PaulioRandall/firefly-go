package token

type Lex struct {
	Token Token
	Value string
	Span  Span
}

func MakeLex(tk Token, val string, span Span) Lex {
	return Lex{
		Token: tk,
		Value: val,
		Span:  span,
	}
}
