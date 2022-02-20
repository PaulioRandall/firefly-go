// Package scanner scans Firefly language tokens from a text source.
//
// To use, call the Begin function with a RuneReader to get the first
// ParseToken function. Invoking it will return a token and the next ParseToken
// function.
package scanner

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// ParseToken is a recursion based tokeniser. It returns the next token and
// a ParseToken function for parsing the next token. On error or while
// obtaining the last token, the returned ParseToken function will be nil.
type ParseToken func() (token.Lexeme, ParseToken, error)

// RuneReader interface is for accessing Go runes from a text source.
type RuneReader interface {

	// More returns true if there are unread runes.
	More() bool

	// Peek returns the next rune without incrementing.
	Peek() (rune, error)

	// Read returns the next rune and increments to the next item.
	Read() (rune, error)
}

// Begin returns a new ParseToken function from which to begin parsing tokens.
// Nil is returned if the supplied reader has already reached the end of
// its stream.
func Begin(r RuneReader) ParseToken {
	if r.More() {
		return scan(r)
	}
	return nil
}

// ScanAll is a convenience function and example for scanning all [remaining]
// tokens from a RuneReader.
func ScanAll(r RuneReader) ([]token.Lexeme, error) {

	var (
		lxs []token.Lexeme
		lx  token.Lexeme
		f   = Begin(r)
		e   error
	)

	for f != nil {
		lx, f, e = f()
		if e != nil {
			return nil, e
		}
		lxs = append(lxs, lx)
	}

	return lxs, nil
}

func scan(r RuneReader) ParseToken {
	return func() (token.Lexeme, ParseToken, error) {

		lx, e := parseToken(r)
		if e != nil {
			return lx, nil, e
		}

		if r.More() {
			return lx, scan(r), nil
		}

		return lx, nil, nil
	}
}

func parseToken(r RuneReader) (token.Lexeme, error) {

	ru, e := r.Read()
	if e != nil {
		return token.Lexeme{}, e
	}

	switch {
	case isNewline(ru):
		return fromRune(token.TK_NEWLINE, ru)

	case isSpace(ru):
		return fromRune(token.TK_SPACE, ru)

	case isLetter(ru):
		return scanWord(r, ru)
	}

	return token.Lexeme{}, newError("unknown token '%v'", string(ru))
}

func scanWord(r RuneReader, first rune) (token.Lexeme, error) {
	word := []rune{first}

	for r.More() {
		next, e := r.Peek()

		if e != nil {
			return token.Lexeme{}, e
		}

		if isSpace(next) {
			goto eval
		}

		r.Read()
		word = append(word, next)
	}

eval:
	w := string(word)
	return evalWord(w)
}

func evalWord(word string) (token.Lexeme, error) {
	switch word {
	case "true", "false":
		return fromStr(token.TK_BOOL, word)
	}

	return token.Lexeme{}, newError("Unknown word '%s'", word)
}

func isNewline(ru rune) bool { return ru == '\n' }
func isSpace(ru rune) bool   { return unicode.IsSpace(ru) && ru != '\n' }
func isLetter(ru rune) bool  { return unicode.IsLetter(ru) }

func fromRune(tk token.Token, ru rune) (token.Lexeme, error) {
	return fromStr(tk, string(ru))
}

func fromStr(tk token.Token, v string) (token.Lexeme, error) {
	lx := token.Lexeme{
		Token: tk,
		Value: v,
	}
	return lx, nil
}

func newError(msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return errors.New(msg)
}
