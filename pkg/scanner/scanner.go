package scanner

import (
	"errors"
	"fmt"
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

	ru, e := sr.Read()
	if e != nil {
		return token.Lexeme{}, e
	}

	var lx token.Lexeme

	switch {
	case isNewline(ru):
		lx = lexeme(token.TokenNewline, ru)
	case isSpace(ru):
		lx = lexeme(token.TokenSpace, ru)
	case isNumber(ru):
		lx = lexeme(token.TokenNumber, ru)
	case ru == '+':
		lx = lexeme(token.TokenAdd, ru)
	case ru == '-':
		lx = lexeme(token.TokenSub, ru)
	case ru == '*':
		lx = lexeme(token.TokenMul, ru)
	case ru == '/':
		lx = lexeme(token.TokenDiv, ru)
	default:
		return lx, newError("Unknown token '%v'", string(ru))
	}

	return lx, nil
}

func lexeme(tk token.Token, ru rune) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: string(ru),
	}
}

func isNewline(ru rune) bool {
	return ru == '\n'
}

func isSpace(ru rune) bool {
	return unicode.IsSpace(ru)
}

func isNumber(ru rune) bool {
	return unicode.IsDigit(ru)
}

func newError(msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return errors.New(msg)
}
