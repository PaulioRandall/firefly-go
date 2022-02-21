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

var empty = token.Lexeme{}

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
		return empty, e
	}

	switch {
	case isNewline(ru):
		return fromRune(token.TK_NEWLINE, ru)

	case isSpace(ru):
		return fromRune(token.TK_SPACE, ru)

	case isDigit(ru):
		return scanNum(r, ru)

	case ru == '"':
		return scanString(r, ru)

	case isChar(ru):
		return scanWord(r, ru)

	default:
		lx, e := scanOperator(r, ru)
		if e != nil {
			return empty, e
		}

		if lx != empty {
			return lx, nil
		}

		return empty, newError("Unknown token '%v'", string(ru))
	}
}

func scanString(r RuneReader, first rune) (token.Lexeme, error) {
	str := []rune{first}
	escape := false

	for r.More() {
		ru, e := r.Peek()

		if e != nil {
			return empty, e
		}

		if isNewline(ru) {
			goto unterminated
		}

		r.Read()
		str = append(str, ru)

		if escape {
			escape = false
			continue
		}

		if ru == '\\' {
			escape = true
			continue
		}

		if ru == '"' {
			return fromStr(token.TK_STRING, string(str))
		}
	}

unterminated:
	return empty, newError("Unterminated string")
}

func scanNum(r RuneReader, first rune) (token.Lexeme, error) {

	// Scan significant part
	num, e := scanInt(r, first)
	if e != nil {
		return empty, e
	}

	if !r.More() {
		return fromStr(token.TK_NUMBER, string(num))
	}

	// Check for decimal point
	ru, e := r.Peek()
	if e != nil {
		return empty, e
	}

	if ru != '.' {
		return fromStr(token.TK_NUMBER, string(num))
	}

	r.Read()
	num = append(num, ru)

	// Scan fractional part
	ru, e = r.Read()
	if e != nil {
		return empty, e
	}

	frac, e := scanInt(r, ru)
	if e != nil {
		return empty, e
	}

	num = append(num, frac...)
	return fromStr(token.TK_NUMBER, string(num))
}

func scanInt(r RuneReader, first rune) ([]rune, error) {
	if !isDigit(first) {
		return nil, newError("Expected digit")
	}

	integer := []rune{first}

	for r.More() {
		ru, e := r.Peek()

		if e != nil {
			return nil, e
		}

		if !isDigit(ru) && ru != '_' {
			break
		}

		r.Read()
		integer = append(integer, ru)
	}

	return integer, nil
}

func scanWord(r RuneReader, first rune) (token.Lexeme, error) {

	if !isChar(first) {
		return empty, newError("Expected letter")
	}

	word := []rune{first}

	for r.More() {
		ru, e := r.Peek()

		if e != nil {
			return empty, e
		}

		if !isChar(ru) {
			break
		}

		r.Read()
		word = append(word, ru)
	}

	w := string(word)
	return evalWord(w)
}

func evalWord(word string) (token.Lexeme, error) {
	switch word {
	case "true", "false":
		return fromStr(token.TK_BOOL, word)
	}

	return fromStr(token.TK_IDENT, word)
}

func scanOperator(r RuneReader, first rune) (token.Lexeme, error) {

	if r.More() {
		lx, e := checkForLongOperator(r, first)
		if e != nil {
			return empty, e
		}

		if lx != empty {
			return lx, nil
		}
	}

	switch first {
	case '+', '-', '*', '/', '%', '=', '<', '>':
		return fromRune(token.TK_OPERATOR, first)
	}

	return empty, nil
}

func checkForLongOperator(r RuneReader, first rune) (token.Lexeme, error) {
	second, e := r.Peek()
	if e != nil {
		return empty, e
	}

	op := string([]rune{first, second})

	switch {
	case op == "==":
	case op == "!=":
	case op == "<=":
	case op == ">=":
	case op == "<<":
	case op == ">>":
	default:
		return empty, nil
	}

	r.Read()
	return fromStr(token.TK_OPERATOR, op)
}

func isNewline(ru rune) bool { return ru == '\n' }
func isSpace(ru rune) bool   { return unicode.IsSpace(ru) && !isNewline(ru) }
func isDigit(ru rune) bool   { return unicode.IsDigit(ru) }
func isChar(ru rune) bool    { return unicode.IsLetter(ru) || ru == '_' }

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

func notImplementedYet() error {
	return newError("Not implemented yet")
}
