package scanner

import (
	"unicode"

	"github.com/PaulioRandall/firefly-go/pkg/err"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

const StringEscape = '\\'
const StringDelim = '"'
const Newline = '\n'

var zeroToken token.Token

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
	ru, e := r.Peek()
	if e != nil {
		return scanTokenFail(r, e)
	}

	sk := sidekick{start: r.Pos()}
	e = scanTokenStartingWith(r, &sk, ru)

	// TODO: Rename 'sidekick' -> 'tokenBuilder'
	// TODO: tokenBuilder shoud have Reader embedded

	if e != nil {
		return scanTokenFail(r, e)
	}

	rng := token.MakeRange(sk.start, r.Pos())
	tk := token.MakeToken(sk.tt, sk.str(), rng)
	return tk, nil
}

func scanTokenStartingWith(r Reader, sk *sidekick, ru rune) error {
	switch {
	case ru == Newline:
		return scanSymbol(r, sk, token.Newline)
	case isSpace(ru):
		return scanSpaces(r, sk)
	case isDigit(ru):
		return scanNumber(r, sk)
	case ru == StringDelim:
		return scanString(r, sk)
	case isWordLetter(ru):
		return scanWord(r, sk)
	default:
		return scanOperator(r, sk)
	}
}

func scanSymbol(r Reader, sk *sidekick, tt token.TokenType) error {
	ru, e := r.Read()
	if e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan %q", tt.String())
	}

	sk.add(ru)
	sk.tt = tt

	return nil
}

func acceptWhile(r Reader, sk *sidekick, f func(ru rune) bool) error {
	added := true
	var e error

	for added {
		added, e = sk.acceptFunc(r, f)

		if e != nil {
			return err.Pos(r.Pos(), e, "Error scanning runes")
		}
	}

	return nil
}

func scanSpaces(r Reader, sk *sidekick) error {
	if e := acceptWhile(r, sk, isSpace); e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan spaces")
	}

	sk.tt = token.Space
	return nil
}

func scanNumber(r Reader, sk *sidekick) error {
	sk.tt = token.Number

	if e := acceptWhile(r, sk, isDigit); e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan significant part of number")
	}

	if hasFractional, e := sk.accept(r, '.'); e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan fractional delimeter of number")
	} else if !hasFractional {
		return nil
	}

	if e := acceptWhile(r, sk, isDigit); e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan fractional part of number")
	}

	return nil
}

func scanString(r Reader, sk *sidekick) error {
	if e := sk.expect(r, StringDelim, "Sanity check!"); e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan initiating string delimiter")
	}

	if e := scanStringBody(r, sk); e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan string body")
	}

	if e := sk.expect(r, StringDelim, "Unterminated string"); e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan terminating string delimiter")
	}

	sk.tt = token.String
	return nil
}

func scanStringBody(r Reader, sk *sidekick) error {
	escaped := false

	e := acceptWhile(r, sk, func(ru rune) bool {
		isDelim := ru == StringDelim

		if !escaped && isDelim {
			return false
		}

		escaped = !escaped && ru == StringEscape
		return true
	})

	if e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan string body")
	}

	if escaped {
		return err.Pos(r.Pos(), nil, "Unterminated string")
	}

	return nil
}

func scanWord(r Reader, sk *sidekick) error {
	if e := acceptWhile(r, sk, isWordLetter); e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan variable or keyword")
	}

	sk.tt = token.IdentifyWordType(string(sk.val))
	return nil
}

func scanOperator(r Reader, sk *sidekick) error {

	var (
		ru1, ru2 rune
		e        error
	)

	ru1, e = r.Read()
	if e != nil {
		return err.Pos(r.Pos(), e, "Failed to scan operator")
	}

	if r.More() {
		ru2, e = r.Peek()
		if e != nil {
			return err.Pos(r.Pos(), e, "Failed to scan operator")
		}
	}

	sk.val = []rune{ru1, ru2}
	sk.tt = token.IdentifyOperatorType(string(sk.val))
	if sk.tt != token.Unknown {
		_, e = r.Read()
		return e
	}

	sk.val = []rune{ru1}
	sk.tt = token.IdentifyOperatorType(string(sk.val))
	if sk.tt != token.Unknown {
		return nil
	}

	if ru2 == rune(0) {
		return err.Pos(r.Pos(), nil, "Unknown symbol %q", ru1)
	}
	return err.Pos(r.Pos(), nil, "Unknown symbol %q", []rune{ru1, ru2})
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

func scanTokenFail(r Reader, e error) (token.Token, error) {
	return zeroToken, err.Pos(r.Pos(), e, "Failed to scan token")
}
