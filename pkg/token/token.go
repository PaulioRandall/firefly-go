package token

// Token is a unit of meaningful assigned to a non-terminal symbol.
type Token int

const (
	TokenUndefined Token = iota
	TokenNewline
	TokenSpace
	TokenParenOpen
	TokenParenClose
	TokenNumber
	TokenAdd
	TokenSub
	TokenMul
	TokenDiv
)

var tokenNames = map[Token]string{
	TokenNewline:    "NEWLINE",
	TokenSpace:      "SPACE",
	TokenParenOpen:  "PAREN_OPEN",
	TokenParenClose: "PAREN_CLOSE",
	TokenNumber:     "NUMBER",
	TokenAdd:        "ADD",
	TokenSub:        "SUBTRACT",
	TokenMul:        "MULTIPLY",
	TokenDiv:        "DIVIDE",
}

// String returns the string representation of the token.
func (tk Token) String() string {
	return tokenNames[tk]
}

func (tk Token) Precedence() int {
	switch tk {
	case TokenParenOpen, TokenParenClose:
		return 3

	case TokenMul, TokenDiv:
		return 2

	case TokenAdd, TokenSub:
		return 1

	default:
		return 0
	}
}
