// Package scanner scans Firefly language tokens from a text source. To use,
// call the Begin function with a RuneReader to get the first ParseToken
// function. Invoking it will return a token and the next ParseToken function.
package scanner

import (
	"errors"
	"fmt"
	"strings"
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
		lxs = []token.Lexeme{}
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

	var lx token.Lexeme
	var e error

	ru, e := r.Read()
	if e != nil {
		return token.Lexeme{}, e
	}

	switch {
	case isNewline(ru):
		lx = lexemeRune(token.TK_NEWLINE, ru)

	case isSpace(ru):
		lx = lexemeRune(token.TK_SPACE, ru)

	case ru == '(':
		lx = lexemeRune(token.TK_PAREN_OPEN, ru)
	case ru == ')':
		lx = lexemeRune(token.TK_PAREN_CLOSE, ru)

	case isNumber(ru):
		lx, e = scanNumber(r, ru)

	case ru == '+':
		lx = lexemeRune(token.TK_ADD, ru)
	case ru == '-':
		lx = lexemeRune(token.TK_SUB, ru)
	case ru == '*':
		lx = lexemeRune(token.TK_MUL, ru)
	case ru == '/':
		lx = lexemeRune(token.TK_DIV, ru)

	default:
		e = newError("Unknown token '%v'", string(ru))
	}

	if e != nil {
		return token.Lexeme{}, e
	}

	return lx, nil
}

func lexemeRune(tk token.Token, ru rune) token.Lexeme {
	return lexemeStr(tk, string(ru))
}

func lexemeStr(tk token.Token, v string) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: v,
	}
}

func scanNumber(r RuneReader, first rune) (token.Lexeme, error) {
	undefined := token.Lexeme{}

	if !r.More() {
		return lexemeRune(token.TK_NUMBER, first), nil
	}

	sb := strings.Builder{}
	sb.WriteRune(first)

	for r.More() {
		ru, e := r.Peek()
		if e != nil {
			return undefined, e
		}

		if !isNumber(ru) {
			break
		}

		_, _ = r.Read()
		sb.WriteRune(ru)
	}

	return lexemeStr(token.TK_NUMBER, sb.String()), nil
}

func isNewline(ru rune) bool { return ru == '\n' }
func isSpace(ru rune) bool   { return unicode.IsSpace(ru) && ru != '\n' }
func isNumber(ru rune) bool  { return unicode.IsDigit(ru) }

func newError(msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return errors.New(msg)
}
