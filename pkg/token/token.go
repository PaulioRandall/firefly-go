// Package token defines the set of tokens that are used to drive parsing.
// A token is a unit of meaning assigned to a non-terminal symbol.
//
// Other shared structures and types are also provided.
package token

// Token is a unit of meaning assigned to a non-terminal symbol. They drive the
// logic for scanning and parsing.
type Token int

const (
	TK_UNDEFINED   Token = iota
	TK_NEWLINE           // \n
	TK_SPACE             // Any whitespace excluding linefeed
	TK_PAREN_OPEN        // (
	TK_PAREN_CLOSE       // )
	TK_NUMBER            // 0, 1, 2, ...
	TK_ADD               // +
	TK_SUB               // -
	TK_MUL               // *
	TK_DIV               // /
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

// Precedence returns the priority of the token so it may be compared with
// others.
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

// IsRedundant returns true if the token is considered redundant to parsing.
func (tk Token) IsRedundant() bool {
	return tk == TK_SPACE
}

// IsCloser returns true if the token closes a block, sub expression, parameter
// set, or condition. Closers usually have a corrisponding opener token.
func (tk Token) IsCloser() bool {
	return tk == TK_PAREN_CLOSE
}

// IsOperator returns true if the token is an operator.
func (tk Token) IsOperator() bool {
	switch tk {
	case TK_PAREN_OPEN, TK_PAREN_CLOSE:
		return true
	case TK_ADD, TK_SUB, TK_MUL, TK_DIV:
		return true
	default:
		return false
	}
}
