package scanner

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// ParseToken is a recursion based tokeniser. It returns the next token and
// another parse function. On error or while obtaining the last token,
// the function will be nil.
type ParseToken func() (token.Lexeme, ParseToken, error)

// ScanAll scans all remaining tokens from the Scroll reader and returns them
// as a slice.
func ScanAll(r token.RuneReader) ([]token.Lexeme, error) {

	var (
		tks = []token.Lexeme{}
		tk  token.Lexeme
		f   = Begin(r)
		e   error
	)

	for f != nil {
		tk, f, e = f()
		if e != nil {
			return nil, e
		}
		tks = append(tks, tk)
	}

	return tks, nil
}

// Begin returns a new function from which to begin parsing tokens.
func Begin(r token.RuneReader) ParseToken {
	if r.More() {
		return scan(r)
	}
	return nil
}

func scan(r token.RuneReader) ParseToken {
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

func parseToken(r token.RuneReader) (token.Lexeme, error) {

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

func scanNumber(r token.RuneReader, first rune) (token.Lexeme, error) {

	if !r.More() {
		return lexemeRune(token.TK_NUMBER, first), nil
	}

	sb := strings.Builder{}
	sb.WriteRune(first)

	for r.More() {
		ru, e := r.Read()
		if e != nil {
			return token.Lexeme{}, e
		}

		if !isNumber(ru) {
			r.PutBack(ru)
			break
		}

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
