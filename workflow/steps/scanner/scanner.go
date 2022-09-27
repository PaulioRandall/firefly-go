// Package scanner converts a string of runes into meaningful tokens
package scanner

import (
	"errors"
	"unicode"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

const (
	StringEscape = '\\'
	StringDelim  = '"'
	Newline      = '\n'
)

var (
	ErrUnknownSymbol      = errors.New("Unknown symbol")
	ErrUnterminatedString = errors.New("Unterminated string")
	ErrMissingFractional  = errors.New("Missing fractional part of number")
	zeroToken             token.Token
)

type RuneReader interface {
	Pos() token.Pos
	More() bool
	Peek() (rune, error)
	Read() (rune, error)
}

type ScanNext func() (tk token.Token, f ScanNext, e error)

func New(rr RuneReader) ScanNext {
	if !rr.More() {
		return nil
	}

	return func() (token.Token, ScanNext, error) {
		tk, e := scanToken(rr)

		if e != nil {
			return zeroToken, nil, e
		}

		return tk, New(rr), nil
	}
}

func scanToken(rr RuneReader) (token.Token, error) {

	var (
		first, second rune
		e             error
		failed        = func(e error) (token.Token, error) {
			return zeroToken, err.AtPos(rr.Pos(), e, "Failed to scan token")
		}
		tb = tokenBuilder{
			rr:    rr,
			start: rr.Pos(),
		}
	)

	if first, e = rr.Read(); e != nil {
		return failed(e)
	}

	if tb.rr.More() {
		if second, e = tb.rr.Peek(); e != nil {
			return failed(e)
		}
	}

	e = scanTokenStartingWith(&tb, first, second)
	if e != nil {
		return failed(e)
	}

	rng := token.MakeRange(tb.start, rr.Pos())
	tk := token.MakeToken(tb.tt, tb.build(), rng)
	return tk, nil
}

func scanTokenStartingWith(tb *tokenBuilder, first, second rune) error {
	switch {
	case isNewline(first):
		tb.tt = token.Newline
		tb.add(first)
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
			return false, tb.err(e, "Failed to accept symbol")
		}
	}

	return accepted, nil
}

func scanSpaces(tb *tokenBuilder, first rune) error {
	tb.add(first)

	if _, e := acceptWhile(tb, isSpace); e != nil {
		return tb.err(e, "Failed to scan spaces")
	}

	tb.tt = token.Space
	return nil
}

func scanComment(tb *tokenBuilder, first rune) error {
	tb.add(first)

	if e := tb.expect(rune('/'), "Sanity check!"); e != nil {
		return e
	}

	if _, e := acceptWhile(tb, isNotNewline); e != nil {
		return tb.err(e, "Failed to scan comment")
	}

	tb.tt = token.Comment
	return nil
}

func scanNumber(tb *tokenBuilder, first rune) error {
	tb.add(first)
	tb.tt = token.Number

	if _, e := acceptWhile(tb, isDigit); e != nil {
		return tb.err(e, "Failed to scan significant part of number")
	}

	hasFractional, e := tb.accept('.')
	if e != nil {
		return tb.err(e, "Failed to scan fractional delimiter in number")
	}

	if hasFractional {
		return scanNumberFraction(tb)
	}

	return nil
}

func scanNumberFraction(tb *tokenBuilder) error {
	hasFractional, e := acceptWhile(tb, isDigit)

	if e != nil {
		return tb.err(e, "Failed to scan fractional part of number")
	}

	if !hasFractional {
		return tb.err(ErrMissingFractional, "Failed to scan fractional part of number")
	}

	return nil
}

func scanString(tb *tokenBuilder, first rune) error {
	tb.add(first)

	if e := scanStringBody(tb); e != nil {
		return tb.err(e, "Failed to scan string body")
	}

	if e := tb.expect(StringDelim, "Unterminated string"); e != nil {
		return tb.err(ErrUnterminatedString, "Failed to scan terminating string delimiter")
	}

	tb.tt = token.String
	return nil
}

func scanStringBody(tb *tokenBuilder) error {
	escaped := false

	_, e := acceptWhile(tb, func(ru rune) bool {
		isDelim := isStringDelim(ru)

		if !escaped && isDelim {
			return false
		}

		escaped = !escaped && ru == StringEscape
		return true
	})

	if e != nil {
		return tb.err(e, "Failed to scan string body")
	}

	if escaped {
		return ErrUnterminatedString
	}

	return nil
}

func scanWord(tb *tokenBuilder, first rune) error {
	tb.add(first)

	if _, e := acceptWhile(tb, isWordLetter); e != nil {
		return tb.err(e, "Failed to scan variable or keyword")
	}

	tb.tt = token.IdentifyWordType(string(tb.val))
	return nil
}

func scanOperator(tb *tokenBuilder, first, second rune) error {

	var scanFail = func(e error) error {
		return tb.err(e, "Failed to scan operator")
	}

	tb.val = []rune{first, second}
	tb.tt = token.IdentifyOperatorType(string(tb.val))
	if tb.tt != token.Unknown {
		_, e := tb.rr.Read()
		if e != nil {
			return scanFail(e)
		}
		return nil
	}

	tb.val = []rune{first}
	tb.tt = token.IdentifyOperatorType(string(tb.val))
	if tb.tt != token.Unknown {
		return nil
	}

	e := unknownSymbol(tb, first, second)
	return scanFail(e)
}

func unknownSymbol(tb *tokenBuilder, first, second rune) error {
	if second == rune(0) {
		return tb.err(ErrUnknownSymbol, "Unknown symbol %q", first)
	}
	return tb.err(ErrUnknownSymbol, "Unknown symbol %q", []rune{first, second})
}

func isNewline(ru rune) bool {
	return ru == '\n'
}

func isNotNewline(ru rune) bool {
	return ru != '\n'
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
