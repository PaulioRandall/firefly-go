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

// begin returns a new ParseToken function.
func begin(sr token.ScrollReader) ParseToken {
	if sr.More() {
		return scan(sr)
	}
	return nil
}

// ScanAll scans all remaining tokens as a slice.
func ScanAll(sr token.ScrollReader) (token.Statement, error) {

	var (
		stmt token.Statement
		tk   token.Lexeme
		f    = begin(sr)
		e    error
	)

	for f != nil {
		tk, f, e = f()
		if e != nil {
			return nil, e
		}
		stmt = append(stmt, tk)
	}

	return stmt, nil
}

func scan(sr token.ScrollReader) ParseToken {
	return func() (token.Lexeme, ParseToken, error) {

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

func parseToken(sr token.ScrollReader) (token.Lexeme, error) {

	var lx token.Lexeme
	var e error

	ru, e := sr.Read()
	if e != nil {
		return token.Lexeme{}, e
	}

	switch {
	case isNewline(ru):
		lx = lexemeRune(token.TokenNewline, ru)
	case isSpace(ru):
		lx = lexemeRune(token.TokenSpace, ru)
	case ru == '(':
		lx = lexemeRune(token.TokenParenOpen, ru)
	case ru == ')':
		lx = lexemeRune(token.TokenParenClose, ru)
	case isNumber(ru):
		lx, e = scanNumber(sr, ru)
	case ru == '+':
		lx = lexemeRune(token.TokenAdd, ru)
	case ru == '-':
		lx = lexemeRune(token.TokenSub, ru)
	case ru == '*':
		lx = lexemeRune(token.TokenMul, ru)
	case ru == '/':
		lx = lexemeRune(token.TokenDiv, ru)
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

func scanNumber(sr token.ScrollReader, first rune) (token.Lexeme, error) {

	if !sr.More() {
		return lexemeRune(token.TokenNumber, first), nil
	}

	sb := strings.Builder{}
	sb.WriteRune(first)

	for sr.More() {
		ru, e := sr.Read()
		if e != nil {
			return token.Lexeme{}, e
		}

		if !isNumber(ru) {
			sr.PutBack(ru)
			break
		}

		sb.WriteRune(ru)
	}

	return lexemeStr(token.TokenNumber, sb.String()), nil
}

func isNewline(ru rune) bool { return ru == '\n' }
func isSpace(ru rune) bool   { return unicode.IsSpace(ru) && ru != '\n' }
func isNumber(ru rune) bool  { return unicode.IsDigit(ru) }

func newError(msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return errors.New(msg)
}
