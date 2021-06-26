package token

// Lexeme is a container for a value and the token it evaluates to. Sometimes
// a lexeme will be referred to as a token. This is because the token drives
// the logic, the value is required only in some cases, e.g. for number tokens.
type Lexeme struct {
	Token
	Value string
}

// Block in the form of a slice of statements.
type Block []Statement

// Statement in the form of a slice of lexemes.
type Statement []Lexeme
