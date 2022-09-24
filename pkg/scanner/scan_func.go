package scanner

import (
	"errors"
	"unicode"

	"github.com/PaulioRandall/firefly-go/pkg/err"
	"github.com/PaulioRandall/firefly-go/pkg/token"
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

type Reader interface {
	Pos() token.Pos
	More() bool
	Peek() (rune, error)
	Read() (rune, error)
}

type ScanFunc func() (tk token.Token, f ScanFunc, e error)

func NewScanFunc(r Reader) ScanFunc {
	if !r.More() {
		return nil
	}

	return func() (token.Token, ScanFunc, error) {
		tk, e := scanToken(r)

		if e != nil {
			return zeroToken, nil, e
		}

		return tk, NewScanFunc(r), nil
	}
}

func scanToken(r Reader) (token.Token, error) {
	failed := func(e error) (token.Token, error) {
		return zeroToken, err.Pos(r.Pos(), e, "Failed to scan token")
	}

	tb := tokenBuilder{
		r:     r,
		start: r.Pos(),
	}

	first, e := r.Read()
	if e != nil {
		return failed(e)
	}

	e = scanTokenStartingWith(&tb, first)
	if e != nil {
		return failed(e)
	}

	rng := token.MakeRange(tb.start, r.Pos())
	tk := token.MakeToken(tb.tt, tb.build(), rng)
	return tk, nil
}

func scanTokenStartingWith(tb *tokenBuilder, first rune) error {
	if first == Newline {
		tb.tt = token.Newline
		tb.add(first)
		return nil
	}

	switch {
	case isSpace(first):
		return scanSpaces(tb, first)
	case isDigit(first):
		return scanNumber(tb, first)
	case first == StringDelim:
		return scanString(tb, first)
	case isWordLetter(first):
		return scanWord(tb, first)
	}

	return scanOperator(tb, first)
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
		isDelim := ru == StringDelim

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

func scanOperator(tb *tokenBuilder, first rune) error {

	var (
		second   rune
		e        error
		scanFail = func(e error) error {
			return tb.err(e, "Failed to scan operator")
		}
	)

	if tb.r.More() {
		second, e = tb.r.Peek()
		if e != nil {
			return scanFail(e)
		}
	}

	tb.val = []rune{first, second}
	tb.tt = token.IdentifyOperatorType(string(tb.val))
	if tb.tt != token.Unknown {
		_, e = tb.r.Read()
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

	e = unknownSymbol(tb, first, second)
	return scanFail(e)
}

func unknownSymbol(tb *tokenBuilder, first, second rune) error {
	if second == rune(0) {
		return tb.err(ErrUnknownSymbol, "Unknown symbol %q", first)
	}
	return tb.err(ErrUnknownSymbol, "Unknown symbol %q", []rune{first, second})
}

func isSpace(ru rune) bool {
	return unicode.IsSpace(ru)
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
