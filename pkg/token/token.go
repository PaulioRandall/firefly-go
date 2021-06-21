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

func (tk Token) IsRedundant() bool {
	return tk == TokenSpace
}

func (tk Token) IsCloser() bool {
	return tk == TokenParenClose
}

func (tk Token) IsOperator() bool {
	switch tk {
	case TokenParenOpen, TokenParenClose:
		fallthrough
	case TokenAdd, TokenSub, TokenMul, TokenDiv:
		return true
	default:
		return false
	}
}
