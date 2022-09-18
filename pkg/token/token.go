package token

type Token struct {
	Type    TokenType
	Value   string
	FilePos Range
}

func MakeToken(tt TokenType, val string, filePos Range) Token {
	return Token{
		Type:    tt,
		Value:   val,
		FilePos: filePos,
	}
}
