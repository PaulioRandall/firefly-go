// Package scanner converts a string of runes into meaningful tokens.
package scanner2

import (
	"io"
	"unicode"

	"github.com/PaulioRandall/go-trackerr"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

var (
	ErrScanning           = trackerr.Checkpoint("Token scanning failed")
	ErrUnknownSymbol      = trackerr.Track("Unknown symbol")
	ErrUnterminatedString = trackerr.Track("Unterminated string")
	ErrMissingFractional  = trackerr.Track("Missing fractional part of number")

	zeroToken token.Token
)

type ReaderOfRunes interface {
	More() bool
	Read() rune
	Putback(rune)
}

type WriterOfTokens interface {
	Write(token.Token) error
}

func Scan(r ReaderOfRunes, w WriterOfTokens) error {
	for r.More() {
		tk, e := scanNext(r)
		if e != nil {
			return ErrScanning.Wrap(e)
		}

		if e := w.Write(tk); e != nil {
			return ErrScanning.CausedBy(e, "Could not write scanned token")
		}
	}

	return nil
}

func scanNext(r ReaderOfRunes) (token.Token, error) {
	first := r.Read()

	switch {
	case isNewline(first):
		return makeToken(token.Newline, str(first)), nil
	case isSpace(first):
		return scanWhitespace(r, first)
	case first == '"':
		return scanString(r, first)
	}

	if r.More() {
		switch second := r.Read(); {
		case first == '/' && second == '/':
			return scanComment(r, first, second)
		case first == '<' && second == '=':
			return makeToken(token.Lte, str(first, second)), nil
		case first == '>' && second == '=':
			return makeToken(token.Gte, str(first, second)), nil
		case first == '=' && second == '=':
			return makeToken(token.Equ, str(first, second)), nil
		case first == '!' && second == '=':
			return makeToken(token.Neq, str(first, second)), nil
		default:
			r.Putback(second)
		}
	}

	switch first {
	case '+':
		return makeToken(token.Add, str(first)), nil
	case '-':
		return makeToken(token.Sub, str(first)), nil
	case '*':
		return makeToken(token.Mul, str(first)), nil
	case '/':
		return makeToken(token.Div, str(first)), nil
	case '%':
		return makeToken(token.Mod, str(first)), nil

	case '<':
		return makeToken(token.Lt, str(first)), nil
	case '>':
		return makeToken(token.Gt, str(first)), nil

	case '=':
		return makeToken(token.Assign, str(first)), nil
	case ':':
		return makeToken(token.Colon, str(first)), nil
	case ';':
		return makeToken(token.Terminator, str(first)), nil
	case ',':
		return makeToken(token.Comma, str(first)), nil
	case '@':
		return makeToken(token.Spell, str(first)), nil

	case '(':
		return makeToken(token.ParenOpen, str(first)), nil
	case ')':
		return makeToken(token.ParenClose, str(first)), nil
	case '{':
		return makeToken(token.BraceOpen, str(first)), nil
	case '}':
		return makeToken(token.BraceClose, str(first)), nil
	case '[':
		return makeToken(token.BracketOpen, str(first)), nil
	case ']':
		return makeToken(token.BracketClose, str(first)), nil
	}

	r.Putback(first)
	return zeroToken, ErrUnknownSymbol.Because("Symbol starting %q could not be resolved", str(first))
}

// Whitespace := ? Any Unicode character from the space category except linefeed ?
func scanWhitespace(r ReaderOfRunes, first rune) (token.Token, error) {
	runes := []rune{first}

	for r.More() {
		ru := r.Read()

		if !isSpace(ru) {
			r.Putback(ru)
			break
		}

		runes = append(runes, ru)
	}

	return makeToken(token.Space, string(runes)), nil
}

// Comment := "//" {Char} Linefeed
// Char    := ? Any Unicode character except linefeed ?
func scanComment(r ReaderOfRunes, first, second rune) (token.Token, error) {
	runes := []rune{first, second}

	for r.More() {
		ru := r.Read()

		if isNewline(ru) {
			r.Putback(ru)
			break
		}

		runes = append(runes, ru)
	}

	return makeToken(token.Comment, string(runes)), nil
}

// String      := '"' {StringChar | '\"'} '"'
// StringChar  := ? Any Unicode character from the letter category except double quote ?
func scanString(r ReaderOfRunes, first rune) (token.Token, error) {
	runes := []rune{first}
	escaped := false
	terminated := false

	for r.More() {
		ru := r.Read()
		runes = append(runes, ru)

		if !escaped && ru == '"' {
			terminated = true
			break
		}

		if isNewline(ru) {
			break
		}

		escaped = !escaped && ru == '\\'
	}

	if !terminated {
		return zeroToken, ErrUnterminatedString.Wrap(io.EOF)
	}

	return makeToken(token.String, string(runes)), nil
}

func makeToken(tt token.TokenType, v string) token.Token {
	return token.Token{
		TokenType: tt,
		Value:     v,
	}
}

func str(runes ...rune) string {
	return string(runes)
}

func isNewline(ru rune) bool {
	return ru == '\n'
}

func isSpace(ru rune) bool {
	return unicode.IsSpace(ru)
}
