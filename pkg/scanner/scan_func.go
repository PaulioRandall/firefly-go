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
	ErrEscapedEndOfString = errors.New("Escaped end of string")
	zeroToken             token.Token
)

type Reader interface {
	Pos() token.Pos
	More() bool
	Peek() (rune, error)
	Read() (rune, error)
}

type ScanFunc func() (tk token.Token, f ScanFunc, e error)

// TODO Test errors

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

	ru, e := r.Peek()
	if e != nil {
		return failed(e)
	}

	tb := tokenBuilder{
		r:     r,
		start: r.Pos(),
	}

	e = scanTokenStartingWith(&tb, ru)
	if e != nil {
		return failed(e)
	}

	rng := token.MakeRange(tb.start, r.Pos())
	tk := token.MakeToken(tb.tt, tb.build(), rng)
	return tk, nil
}

func scanTokenStartingWith(tb *tokenBuilder, ru rune) error {
	switch {
	case ru == Newline:
		return scanSymbol(tb, token.Newline)
	case isSpace(ru):
		return scanSpaces(tb)
	case isDigit(ru):
		return scanNumber(tb)
	case ru == StringDelim:
		return scanString(tb)
	case isWordLetter(ru):
		return scanWord(tb)
	default:
		return scanOperator(tb)
	}
}

func scanSymbol(tb *tokenBuilder, tt token.TokenType) error {
	if e := tb.any(); e != nil {
		return tb.err(e, "Failed to scan symbol")
	}
	tb.tt = tt
	return nil
}

func acceptWhile(tb *tokenBuilder, f func(ru rune) bool) error {
	added := true
	var e error

	for added {
		added, e = tb.acceptFunc(f)

		if e != nil {
			return tb.err(e, "Scanning error")
		}
	}

	return nil
}

func scanSpaces(tb *tokenBuilder) error {
	if e := acceptWhile(tb, isSpace); e != nil {
		return tb.err(e, "Failed to scan spaces")
	}

	tb.tt = token.Space
	return nil
}

func scanNumber(tb *tokenBuilder) error {
	tb.tt = token.Number

	if e := acceptWhile(tb, isDigit); e != nil {
		return tb.err(e, "Failed to scan significant part of number")
	}

	if hasFractional, e := tb.accept('.'); e != nil {
		return tb.err(e, "Failed to scan fractional delimiter in number")
	} else if !hasFractional {
		return nil
	}

	if e := acceptWhile(tb, isDigit); e != nil {
		return tb.err(e, "Failed to scan fractional part of number")
	}

	return nil
}

func scanString(tb *tokenBuilder) error {
	if e := tb.expect(StringDelim, "Sanity check!"); e != nil {
		return tb.err(e, "Failed to scan initial string delimiter")
	}

	if e := scanStringBody(tb); e != nil {
		return tb.err(e, "Failed to scan string body")
	}

	if e := tb.expect(StringDelim, "Unterminated string"); e != nil {
		return tb.err(e, "Failed to scan terminating string delimiter")
	}

	tb.tt = token.String
	return nil
}

func scanStringBody(tb *tokenBuilder) error {
	escaped := false

	e := acceptWhile(tb, func(ru rune) bool {
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
		return tb.err(ErrEscapedEndOfString, "Unterminated string")
	}

	return nil
}

func scanWord(tb *tokenBuilder) error {
	if e := acceptWhile(tb, isWordLetter); e != nil {
		return tb.err(e, "Failed to scan variable or keyword")
	}

	tb.tt = token.IdentifyWordType(string(tb.val))
	return nil
}

func scanOperator(tb *tokenBuilder) error {

	var (
		ru1, ru2 rune
		e        error
		scanFail = func(e error) error {
			return tb.err(e, "Failed to scan operator")
		}
	)

	ru1, e = tb.r.Read()
	if e != nil {
		return scanFail(e)
	}

	if tb.r.More() {
		ru2, e = tb.r.Peek()
		if e != nil {
			return scanFail(e)
		}
	}

	tb.val = []rune{ru1, ru2}
	tb.tt = token.IdentifyOperatorType(string(tb.val))
	if tb.tt != token.Unknown {
		_, e = tb.r.Read()
		if e != nil {
			return scanFail(e)
		}
		return nil
	}

	tb.val = []rune{ru1}
	tb.tt = token.IdentifyOperatorType(string(tb.val))
	if tb.tt != token.Unknown {
		return nil
	}

	e = unknownSymbol(tb, ru1, ru2)
	return scanFail(e)
}

func unknownSymbol(tb *tokenBuilder, ru1, ru2 rune) error {
	if ru2 == rune(0) {
		return tb.err(ErrUnknownSymbol, "Unknown symbol %q", ru1)
	}
	return tb.err(ErrUnknownSymbol, "Unknown symbol %q", []rune{ru1, ru2})
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
