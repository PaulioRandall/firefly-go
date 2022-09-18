package token

type Token struct {
	Type  TokenType
	Value string
	Span  Span
}

func MakeToken(tt TokenType, val string, span Span) Token {
	return Token{
		Type:  tt,
		Value: val,
		Span:  span,
	}
}
