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

type TokenGenerator func(TokenType, string) Token

func NewTokenGenerator() TokenGenerator {
	cursor := Pos{}

	return func(tt TokenType, v string) Token {
		start := cursor
		cursor.IncString(v)
		rng := MakeRange(start, cursor)
		return MakeToken(tt, v, rng)
	}
}
