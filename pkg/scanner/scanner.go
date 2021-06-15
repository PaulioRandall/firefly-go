package scanner

import (
	"errors"
	"fmt"
	"unicode"
)

// ScrollReader is the interface for accessing Go runes from a text source.
type ScrollReader interface {

	// More returns true if there are unread runes from the text source.
	More() bool

	// Read returns the next rune in the scroll and moves the read head to the
	// next rune.
	Read() (rune, error)

	// PutBack puts a rune back into the scoll reader so it becomes the next
	// rune to be read.
	PutBack(rune) error
}

// ParseToken is a recursion based tokeniser. It returns the next token and
// another parse function. On error or while obtaining the last token,
// the function will be nil.
type ParseToken func() (Lexeme, ParseToken, error)

// New returns a new ParseToken function.
func New(sr ScrollReader) ParseToken {
	if sr.More() {
		return scan(sr)
	}
	return nil
}

// ScanAll scans all remaining tokens as a slice.
func ScanAll(sr ScrollReader) ([]Lexeme, error) {

	var (
		result []Lexeme
		tk     Lexeme
		f      = New(sr)
		e      error
	)

	for f != nil {
		tk, f, e = f()
		if e != nil {
			return nil, e
		}
		result = append(result, tk)
	}

	return result, nil
}

func scan(sr ScrollReader) ParseToken {
	return func() (Lexeme, ParseToken, error) {

		lx, e := parseToken(sr)
		if e != nil {
			return lx, nil, e
		}

		if sr.More() {
			return lx, scan(sr), nil
		}

		return lx, nil, nil
	}
}

func parseToken(sr ScrollReader) (Lexeme, error) {

	ru, e := sr.Read()
	if e != nil {
		return Lexeme{}, e
	}

	var lx Lexeme

	switch {
	case unicode.IsSpace(ru):
		lx = Lexeme{TokenSpace, string(ru)}
	case unicode.IsDigit(ru):
		lx = Lexeme{TokenNumber, string(ru)}
	case ru == '+':
		lx = Lexeme{TokenOperator, string(ru)}
	default:
		return lx, newError("Unknown token '%v'", string(ru))
	}

	return lx, nil
}

func newError(msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return errors.New(msg)
}
