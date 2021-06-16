package token

// Token is a unit of meaningful assigned to a non-terminal symbol
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

// Name returns the string representation of the token
func (tk Token) Name() string {
	return tokenNames[tk]
}

// Lexeme is a value with associated token
type Lexeme struct {
	Token
	Value string
}
