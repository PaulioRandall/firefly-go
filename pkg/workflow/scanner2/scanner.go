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
	case isLetter(first):
		return scanWord(r, first)
	case isDigit(first):
		return scanNumber(r, first)
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
		case first == '&' && second == '&':
			return makeToken(token.And, str(first, second)), nil
		case first == '|' && second == '|':
			return makeToken(token.Or, str(first, second)), nil
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
	return zeroToken, ErrUnknownSymbol.Because("Symbol could not be resolved")
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

// Def   := "def"
// If    := "if"
// For   := "for"
// In    := "in"
// Watch := "watch"
// When  := "when"
// Is    := "is"
// Func  := "F"
// Proc  := "P"
// End   := "end"
//
// Bool := "true" | "false"
//
// Ident     := IdentChar {IdentChar}
// IdentChar := "_" | ? Any Unicode character from the letter category ?
func scanWord(r ReaderOfRunes, first rune) (token.Token, error) {
	runes := []rune{first}

	for r.More() {
		ru := r.Read()

		if !isLetter(ru) {
			r.Putback(ru)
			break
		}

		runes = append(runes, ru)
	}

	tk := token.Token{
		Value: string(runes),
	}

	switch tk.Value {
	case "def":
		tk.TokenType = token.Def
	case "if":
		tk.TokenType = token.If
	case "for":
		tk.TokenType = token.For
	case "in":
		tk.TokenType = token.In
	case "watch":
		tk.TokenType = token.Watch
	case "when":
		tk.TokenType = token.When
	case "is":
		tk.TokenType = token.Is
	case "F":
		tk.TokenType = token.Func
	case "P":
		tk.TokenType = token.Proc
	case "end":
		tk.TokenType = token.End
	case "true", "false":
		tk.TokenType = token.Bool
	default:
		tk.TokenType = token.Ident
	}

	return tk, nil
}

// Number  := Integer ["." Integer]
// Int   := Digit {Digit}
// Digit := '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'
func scanNumber(r ReaderOfRunes, first rune) (token.Token, error) {
	r.Putback(first)
	runes := scanInt(r)

	if !r.More() {
		return makeToken(token.Number, string(runes)), nil
	}

	if ru := r.Read(); ru == '.' {
		runes = append(runes, ru)
	} else {
		r.Putback(ru)
		return makeToken(token.Number, string(runes)), nil
	}

	if fractional := scanInt(r); len(fractional) > 0 {
		runes = append(runes, fractional...)
	} else {
		return zeroToken, ErrMissingFractional
	}

	return makeToken(token.Number, string(runes)), nil
}

func scanInt(r ReaderOfRunes) []rune {
	var runes []rune

	for r.More() {
		ru := r.Read()

		if !isDigit(ru) {
			r.Putback(ru)
			break
		}

		runes = append(runes, ru)
	}

	return runes
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

func isLetter(ru rune) bool {
	return unicode.IsLetter(ru) || ru == '_'
}

func isDigit(ru rune) bool {
	switch ru {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}
