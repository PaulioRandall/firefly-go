// Package token defines the set of tokens that are used to drive parsing.
// A token is a unit of meaning assigned to a non-terminal symbol.
package token

// Token is a unit of meaning assigned to a non-terminal symbol. They drive the
// logic for scanning and parsing.
type Token int

const (
	TK_UNDEFINED Token = iota
	TK_NEWLINE         // \n
	TK_SPACE           // Any whitespace excluding linefeed
	TK_BOOL            // 'true' or 'false'
	TK_NUMBER          // 123.456
	TK_STRING          // ""
	TK_IDENT           // Identifier (variable name)
	TK_OPERATOR        // +, >=, <<, etc
)

var tokenNames = map[Token]string{
	TK_NEWLINE:  "NEWLINE",
	TK_SPACE:    "SPACE",
	TK_BOOL:     "BOOL",
	TK_NUMBER:   "NUMBER",
	TK_STRING:   "STRING",
	TK_IDENT:    "IDENT",
	TK_OPERATOR: "OPERATOR",
}

// String returns the string representation of the token.
func (tk Token) String() string {
	return tokenNames[tk]
}

// Precedence returns the priority of the token within an expression.
func (tk Token) Precedence() int {
	switch tk {
	default:
		return 0
	}
}

// IsRedundant returns true if the token is considered redundant to parsing.
func (tk Token) IsRedundant() bool {
	return tk == TK_SPACE
}
