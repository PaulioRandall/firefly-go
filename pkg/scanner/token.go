package scanner

type Token int

const (
	TokenUndefined Token = iota
	TokenSpace
	TokenNumber
	TokenOperator
)

var tokenNames = map[Token]string{
	TokenSpace:    "SPACE",
	TokenNumber:   "NUMBER",
	TokenOperator: "OPERATOR",
}

func (tk Token) Name() string {
	return tokenNames[tk]
}

type Lexeme struct {
	Token
	Value string
}
