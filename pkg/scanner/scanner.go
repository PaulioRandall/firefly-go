package scanner

import (
	"unicode"

	"github.com/PaulioRandall/firefly-go/pkg/err"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

var zeroToken token.Token

type Reader interface {
	Pos() token.Pos
	More() bool
	Peek() (rune, error)
	Read() (rune, error)
}

type ScanFunc func() (tk token.Token, f ScanFunc, e error)

func ScanAll(r Reader) ([]token.Token, error) {
	var (
		tk  token.Token
		tks []token.Token
		sc  = NewScanFunc(r)
		e   error
	)

	for sc != nil {
		tk, sc, e = sc()

		if e != nil {
			return nil, err.Pos(r.Pos(), e, "Failed to scan all tokens")
		}

		tks = append(tks, tk)
	}

	return tks, nil
}

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
	var (
		tt    token.TokenType
		value string
		start = r.Pos()
	)

	ru, e := r.Peek()
	if e != nil {
		return scanTokenFail(r, e)
	}

	switch {
	case isWordLetter(ru):
		value, tt, e = scanWord(r)
	default:
		return unknownSymbol(r, ru)
	}

	if e != nil {
		return scanTokenFail(r, e)
	}

	tk := token.MakeToken(
		tt,
		value,
		token.MakeRange(start, r.Pos()),
	)
	return tk, nil
}

func scanWord(r Reader) (string, token.TokenType, error) {
	var runes []rune

	for r.More() {
		ru, e := r.Peek()
		if e != nil {
			return scanWordFail(r, e)
		}

		if !isWordLetter(ru) {
			break
		}

		_, e = r.Read()
		if e != nil {
			return scanWordFail(r, e)
		}

		runes = append(runes, ru)
	}

	word := string(runes)
	return word, token.IdentifyWordType(word), nil
}

func isWordLetter(ru rune) bool {
	return unicode.IsLetter(ru) || ru == '_'
}

func scanWordFail(r Reader, e error) (string, token.TokenType, error) {
	return "", token.Unknown, err.Pos(r.Pos(), e, "Failed to scan word")
}

func scanTokenFail(r Reader, e error) (token.Token, error) {
	return zeroToken, err.Pos(r.Pos(), e, "Failed to scan token")
}

func unknownSymbol(r Reader, sym rune) (token.Token, error) {
	return zeroToken, err.Pos(r.Pos(), nil, "Unknown symbol %q", sym)
}
