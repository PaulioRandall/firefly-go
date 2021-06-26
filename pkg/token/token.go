package token

// Token is a unit of meaningful assigned to a non-terminal symbol.
type Token int

const (
	TK_UNDEFINED Token = iota
	TK_NEWLINE
	TK_SPACE
	TK_PAREN_OPEN
	TK_PAREN_CLOSE
	TK_NUMBER
	TK_ADD
	TK_SUB
	TK_MUL
	TK_DIV
)

var tokenNames = map[Token]string{
	TK_NEWLINE:     "NEWLINE",
	TK_SPACE:       "SPACE",
	TK_PAREN_OPEN:  "PAREN_OPEN",
	TK_PAREN_CLOSE: "PAREN_CLOSE",
	TK_NUMBER:      "NUMBER",
	TK_ADD:         "ADD",
	TK_SUB:         "SUBTRACT",
	TK_MUL:         "MULTIPLY",
	TK_DIV:         "DIVIDE",
}

// String returns the string representation of the token.
func (tk Token) String() string {
	return tokenNames[tk]
}

func (tk Token) Precedence() int {
	switch tk {
	case TK_PAREN_OPEN, TK_PAREN_CLOSE:
		return 3

	case TK_MUL, TK_DIV:
		return 2

	case TK_ADD, TK_SUB:
		return 1

	default:
		return 0
	}
}

func (tk Token) IsRedundant() bool {
	return tk == TK_SPACE
}

func (tk Token) IsCloser() bool {
	return tk == TK_PAREN_CLOSE
}

func (tk Token) IsOperator() bool {
	switch tk {
	case TK_PAREN_OPEN, TK_PAREN_CLOSE:
		fallthrough
	case TK_ADD, TK_SUB, TK_MUL, TK_DIV:
		return true
	default:
		return false
	}
}
