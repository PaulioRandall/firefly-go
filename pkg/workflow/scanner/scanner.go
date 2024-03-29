// Package scanner converts a string of runes into meaningful tokens.
package scanner

import (
	"unicode"

	"github.com/PaulioRandall/go-trackerr"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

const (
	StringEscape = '\\'
	StringDelim  = '"'
	Newline      = '\n'
	Terminator   = ';'
)

var (
	ErrScanning           = trackerr.Checkpoint("Token scanning failed")
	ErrUnknownSymbol      = trackerr.Track("Unknown symbol")
	ErrUnterminatedString = trackerr.Track("Unterminated string")
	ErrMissingFractional  = trackerr.Track("Missing fractional part of number")
	zeroToken             token.Token
)

type ReaderOfRunes inout.Reader[rune]
type WriterOfTokens inout.Writer[token.Token]

func Scan(r ReaderOfRunes, w WriterOfTokens) error {
	tb := newTokenBuilder(r)

	for tb.More() {
		if e := scanNext(&tb); e != nil {
			return ErrScanning.Wrap(e)
		}

		tk := tb.build()
		if e := w.Write(tk); e != nil {
			return ErrScanning.CausedBy(e, "Could not write scanned token")
		}
	}

	return nil
}

func scanNext(tb *tokenBuilder) error {

	first, e := tb.Read()
	if e != nil {
		return e
	}

	second, e := tb.Peek()
	if e != nil && trackerr.Is(e, inout.EOF) {
		second, e = rune(0), nil
	} else if e != nil {
		return e
	}

	if e = scanToken(tb, first, second); e != nil {
		return e
	}

	return nil
}

func scanToken(tb *tokenBuilder, first, second rune) error {
	switch {
	case isNewline(first):
		tb.tt = token.Newline
		return nil

	case isTerminator(first):
		tb.tt = token.Terminator
		return nil

	case isSpace(first):
		return scanSpaces(tb, first)

	case isCommentPrefix(first, second):
		return scanComment(tb, first)

	case isDigit(first):
		return scanNumber(tb, first)

	case isStringDelim(first):
		return scanString(tb, first)

	case isWordLetter(first):
		return scanWord(tb, first)
	}

	return scanOperator(tb, first, second)
}

func acceptWhile(tb *tokenBuilder, f func(ru rune) bool) (bool, error) {
	accepted := false

	for hasNext, e := tb.acceptFunc(f); hasNext; hasNext, e = tb.acceptFunc(f) {
		accepted = true

		if e != nil {
			return false, trackerr.Wrap(e, "Failed to accept symbol")
		}
	}

	return accepted, nil
}

// Whitespace := ? Any Unicode character from the space category except linefeed ?
func scanSpaces(tb *tokenBuilder, first rune) error {
	if _, e := acceptWhile(tb, isSpace); e != nil {
		return trackerr.Wrap(e, "Failed to scan spaces")
	}

	tb.tt = token.Space
	return nil
}

// Comment := "//" {Char} Linefeed
// Char    := ? Any Unicode character except linefeed ?
func scanComment(tb *tokenBuilder, first rune) error {
	if e := tb.expect(rune('/'), "Sanity check!"); e != nil {
		return trackerr.Wrap(e, "Failed to scan comment")
	}

	if _, e := acceptWhile(tb, isNotNewline); e != nil {
		return trackerr.Wrap(e, "Failed to scan comment")
	}

	tb.tt = token.Comment
	return nil
}

// Number  := Integer ["." Integer]
// Integer := Digit {Digit}
// Digit   := '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'
func scanNumber(tb *tokenBuilder, first rune) error {
	tb.tt = token.Number

	if _, e := acceptWhile(tb, isDigit); e != nil {
		return trackerr.Wrap(e, "Failed to scan significant part of number")
	}

	hasFractional, e := tb.accept('.')
	if e != nil {
		return trackerr.Wrap(e, "Failed to scan fractional delimiter in number")
	}

	if hasFractional {
		return scanNumberFraction(tb)
	}

	return nil
}

func scanNumberFraction(tb *tokenBuilder) error {
	hasFractional, e := acceptWhile(tb, isDigit)

	if e != nil {
		return trackerr.Wrap(e, "Failed to scan fractional part of number")
	}

	if !hasFractional {
		return ErrMissingFractional.Because("Expected fractional digits")
	}

	return nil
}

// String      := '"' {StringChar | '\"'} '"'
// StringChar  := ? Any Unicode character from the letter category except double quote ?
func scanString(tb *tokenBuilder, first rune) error {
	if e := scanStringBody(tb); e != nil {
		return trackerr.Wrap(e, "Failed to scan string body")
	}

	if e := tb.expect(StringDelim, "Unterminated string"); e != nil {
		return ErrUnterminatedString.CausedBy(e, "Failed to scan terminating delimiter")
	}

	tb.tt = token.String
	return nil
}

func scanStringBody(tb *tokenBuilder) error {
	escaped := false
	terminated := false

	_, e := acceptWhile(tb, func(ru rune) bool {
		isDelim := isStringDelim(ru)

		if !escaped && isDelim {
			terminated = true
			return false
		}

		escaped = !escaped && ru == StringEscape
		return true
	})

	if e != nil {
		return trackerr.Wrap(e, "Failed to scan string body")
	}

	if !terminated {
		return ErrUnterminatedString.Because("EOF before string terminated")
	}

	return nil
}

// Def       := "def"
// If        := "if"
// For       := "for"
// In        := "in"
// Watch     := "watch"
// When      := "when"
// Is        := "is"
// E         := "E"
// P         := "P"
// End       := "end"
// Bool      := "true" | "false"
// Ident     := IdentChar {IdentChar}
// IdentChar := "_" | ? Any Unicode character from the letter category ?
func scanWord(tb *tokenBuilder, first rune) error {
	if _, e := acceptWhile(tb, isWordLetter); e != nil {
		return trackerr.Wrap(e, "Failed to scan variable or keyword")
	}

	switch tb.String() {
	case "def":
		tb.tt = token.Def
	case "if":
		tb.tt = token.If
	case "for":
		tb.tt = token.For
	case "in":
		tb.tt = token.In
	case "watch":
		tb.tt = token.Watch
	case "when":
		tb.tt = token.When
	case "is":
		tb.tt = token.Is
	case "F":
		tb.tt = token.Func
	case "P":
		tb.tt = token.Proc
	case "end":
		tb.tt = token.End
	case "true", "false":
		tb.tt = token.Bool
	default:
		tb.tt = token.Ident
	}

	return nil
}

// Terminator   := ";"
// Assign       := "="
// Comma        := ","
// Colon        := ":"
// Spell        := "@"
//
// Add          := "+"
// Sub          := "-"
// Mul          := "*"
// Div          := "/"
// Mod          := "%"
//
// LT           := "<"
// GT           := ">"
// LTE          := "<="
// GTE          := ">="
// EQU          := "=="
// NEQ          := "!="
//
// ParenOpen    := "("
// ParenClose   := ")"
// BraceOpen    := "{"
// BraceClose   := "}"
// BracketOpen  := "["
// BracketClose := "]"
func scanOperator(tb *tokenBuilder, first, second rune) error {

	one := func(tb *tokenBuilder, tt token.TokenType) error {
		tb.tt = tt
		return nil
	}

	two := func(tb *tokenBuilder, tt token.TokenType) error {
		tb.tt = tt

		if _, e := tb.accept(second); e != nil {
			return trackerr.Wrap(e, "Failed to scan operator")
		}

		return nil
	}

	switch {
	case first == '+':
		return one(tb, token.Add)

	case first == '-':
		return one(tb, token.Sub)

	case first == '*':
		return one(tb, token.Mul)

	case first == '/':
		return one(tb, token.Div)

	case first == '%':
		return one(tb, token.Mod)

	case first == '<' && second == '=':
		return two(tb, token.Lte)

	case first == '>' && second == '=':
		return two(tb, token.Gte)

	case first == '<':
		return one(tb, token.Lt)

	case first == '>':
		return one(tb, token.Gt)

	case first == '=' && second == '=':
		return two(tb, token.Equ)

	case first == '!' && second == '=':
		return two(tb, token.Neq)

	case first == '=':
		return one(tb, token.Assign)

	case first == '&' && second == '&':
		return two(tb, token.And)

	case first == '|' && second == '|':
		return two(tb, token.Or)

	case first == ':':
		return one(tb, token.Colon)

	case first == ';':
		return one(tb, token.Terminator)

	case first == ',':
		return one(tb, token.Comma)

	case first == '@':
		return one(tb, token.Spell)

	case first == '(':
		return one(tb, token.ParenOpen)

	case first == ')':
		return one(tb, token.ParenClose)

	case first == '{':
		return one(tb, token.BraceOpen)

	case first == '}':
		return one(tb, token.BraceClose)

	case first == '[':
		return one(tb, token.BracketOpen)

	case first == ']':
		return one(tb, token.BracketClose)

	default:
		e := unknownSymbol(tb, first, second)
		return trackerr.Wrap(e, "Failed to scan operator")
	}
}

func unknownSymbol(tb *tokenBuilder, first, second rune) error {
	if second == rune(0) {
		return ErrUnknownSymbol.Because("Unknown symbol %q", first)
	}
	return ErrUnknownSymbol.Because("Unknown symbol %q", []rune{first, second})
}

func isNewline(ru rune) bool {
	return ru == '\n'
}

func isNotNewline(ru rune) bool {
	return ru != '\n'
}

func isTerminator(ru rune) bool {
	return ru == ';'
}

func isSpace(ru rune) bool {
	return unicode.IsSpace(ru)
}

func isCommentPrefix(first, second rune) bool {
	return first == '/' && second == '/'
}

func isStringDelim(ru rune) bool {
	return ru == StringDelim
}

func isWordLetter(ru rune) bool {
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
