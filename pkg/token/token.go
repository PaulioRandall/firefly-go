package token

// Token is a unit of meaningful assigned to a non-terminal symbol.
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

// String returns the string representation of the token.
func (tk Token) String() string {
	return tokenNames[tk]
}

/*
func (tk Token) Precedence() int {
	switch {
	case
	}
}
*/
