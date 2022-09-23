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
	prev := Range{}

	return func(tt TokenType, v string) Token {
		prev = prev.IncString(v)
		return MakeToken(tt, v, prev)
	}
}
