// Package scanner converts a string of runes into meaningful tokens
package scanner

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/PaulioRandall/firefly-go/utilities/inout"
	"github.com/PaulioRandall/firefly-go/utilities/pos"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

const (
	StringEscape = '\\'
	StringDelim  = '"'
	Newline      = '\n'
	Terminator   = ';'
)

var (
	ErrUnknownSymbol      = errors.New("Unknown symbol")
	ErrUnterminatedString = errors.New("Unterminated string")
	ErrMissingFractional  = errors.New("Missing fractional part of number")
	zeroToken             token.Token
)

type ReaderOfRunes inout.Reader[rune]
type TokenWriter inout.Writer[token.Token]

func Scan(r ReaderOfRunes, w TokenWriter) error {
	tb := newTokenBuilder(r)

	for tb.More() {
		if e := scanNext(&tb); e != nil {
			return fmt.Errorf("Failed to scan tokens: %w", e)
		}

		tk := tb.build()
		if e := w.Write(tk); e != nil {
			return fmt.Errorf("Failed to scan tokens: %w", e)
		}
	}

	return nil
}

func scanNext(tb *tokenBuilder) error {

	failed := func(e error) error {
		return tb.err(e, "Failed to scan token")
	}

	first, e := tb.Read()
	if e != nil {
		return failed(e)
	}

	second, e := tb.Peek()
	if e == inout.EOF {
		second, e = rune(0), nil
	} else if e != nil {
		return failed(e)
	}

	if e = scanToken(tb, first, second); e != nil {
		return failed(e)
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
			return false, tb.err(e, "Failed to accept symbol")
		}
	}

	return accepted, nil
}

func scanSpaces(tb *tokenBuilder, first rune) error {
	if _, e := acceptWhile(tb, isSpace); e != nil {
		return tb.err(e, "Failed to scan spaces")
	}

	tb.tt = token.Space
	return nil
}

func scanComment(tb *tokenBuilder, first rune) error {
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
	if _, e := acceptWhile(tb, isWordLetter); e != nil {
		return tb.err(e, "Failed to scan variable or keyword")
	}

	tb.tt = token.IdentifyWordType(tb.String())
	return nil
}

func scanOperator(tb *tokenBuilder, first, second rune) error {

	var scanFail = func(e error) error {
		return tb.err(e, "Failed to scan operator")
	}

	ok, e := tryScanTwoRuneOperator(tb, first, second)
	if e != nil {
		return scanFail(e)
	}

	if ok {
		return nil
	}

	if tryScanSingleRuneOperator(tb, first) {
		return nil
	}

	e = unknownSymbol(tb, first, second)
	return scanFail(e)
}

func tryScanTwoRuneOperator(tb *tokenBuilder, first, second rune) (bool, error) {
	val := []rune{first, second}
	tt := token.IdentifyOperatorType(string(val))

	if tt == token.Unknown {
		return false, nil
	}

	if ok, e := tb.accept(second); !ok || e != nil {
		return false, tb.err(e, "Failed to scan second symbol in operator")
	}

	tb.tt = tt
	return true, nil
}

func tryScanSingleRuneOperator(tb *tokenBuilder, first rune) bool {
	val := []rune{first}
	tt := token.IdentifyOperatorType(string(val))

	if tt == token.Unknown {
		return false
	}

	tb.tt = tt
	return true
}

func unknownSymbol(tb *tokenBuilder, first, second rune) error {
	if second == rune(0) {
		return pos.ErrorFor(tb.start, ErrUnknownSymbol, "Unknown symbol %q", first)
	}
	return pos.ErrorFor(tb.start, ErrUnknownSymbol, "Unknown symbol %q", []rune{first, second})
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
