package scanner

type Token int

const (
	TokenUndefined Token = iota
	TokenNewline
	TokenSpace
	TokenNumber
	TokenOperator
)

var tokenNames = map[Token]string{
	TokenNewline:  "NEWLINE",
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
