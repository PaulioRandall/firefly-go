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
		val   string
		start = r.Pos()
	)

	ru, e := r.Peek()
	if e != nil {
		return scanTokenFail(r, e)
	}

	switch {
	case isWordLetter(ru):
		val, tt, e = scanWord(r)
	default:
		val, tt, e = scanOperator(r)
	}

	if e != nil {
		return scanTokenFail(r, e)
	}

	rng := token.MakeRange(start, r.Pos())
	tk := token.MakeToken(tt, val, rng)
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

func scanOperator(r Reader) (string, token.TokenType, error) {

	var (
		ru1, ru2 rune
		e        error
		tt       token.TokenType
		val      string
	)

	ru1, e = r.Read()
	if e != nil {
		return scanOperatorFail(r, e)
	}

	if r.More() {
		ru2, e = r.Peek()
		if e != nil {
			return scanOperatorFail(r, e)
		}
	}

	val = string([]rune{ru1, ru2})
	tt = token.IdentifyOperatorType(val)
	if tt != token.Unknown {
		_, e = r.Read()
		return val, tt, e
	}

	val = string([]rune{ru1})
	tt = token.IdentifyOperatorType(val)
	if tt != token.Unknown {
		return val, tt, nil
	}

	return unknownSymbol(r, ru1, ru2)
}

func isWordLetter(ru rune) bool {
	return unicode.IsLetter(ru) || ru == '_'
}

func scanTokenFail(r Reader, e error) (token.Token, error) {
	return zeroToken, err.Pos(r.Pos(), e, "Failed to scan token")
}

func scanWordFail(r Reader, e error) (string, token.TokenType, error) {
	return "", token.Unknown, err.Pos(r.Pos(), e, "Failed to scan word")
}

func scanOperatorFail(r Reader, e error) (string, token.TokenType, error) {
	return "", token.Unknown, err.Pos(r.Pos(), e, "Failed to scan operator")
}

func unknownSymbol(r Reader, sym1, sym2 rune) (string, token.TokenType, error) {
	if sym2 == rune(0) {
		return "", token.Unknown, err.Pos(r.Pos(), nil, "Unknown symbol %q", sym1)
	}
	return "", token.Unknown, err.Pos(r.Pos(), nil, "Unknown symbol %q", []rune{sym1, sym2})
}
