package token

type Lex struct {
	tk   Token
	val  string
	span Span
}

func MakeLex(tk Token, val string, span Span) Lex {
	return Lex{
		tk:   tk,
		val:  val,
		span: span,
	}
}
