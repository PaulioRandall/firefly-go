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
	case ru == Newline:
		val, tt, e = scanNewline(r)
	case isSpace(ru):
		val, tt, e = scanSpace(r)
	case isDigit(ru):
		val, tt, e = scanNumber(r)
	case ru == StringDelim:
		val, tt, e = scanString(r)
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

func scanNewline(r Reader) (string, token.TokenType, error) {
	ru, e := r.Read()
	if e != nil {
		return scanNewlineFail(r, e)
	}

	return string(ru), token.Newline, nil
}

func scanSpace(r Reader) (string, token.TokenType, error) {
	var spaces []rune

	for r.More() {
		ru, e := r.Peek()
		if e != nil {
			return scanSpaceFail(r, e)
		}

		if ru == Newline || !isSpace(ru) {
			break
		}

		_, e = r.Read()
		if e != nil {
			return scanSpaceFail(r, e)
		}

		spaces = append(spaces, ru)
	}

	return string(spaces), token.Space, nil
}

func scanNumber(r Reader) (string, token.TokenType, error) {
	sig, e := scanInt(r)
	if e != nil {
		return scanNumberFail(r, e)
	}

	if !r.More() {
		return sig, token.Number, nil
	}

	dot, e := r.Peek()
	if e != nil {
		return scanNumberFail(r, e)
	}

	if dot != '.' {
		return sig, token.Number, nil
	}

	if _, e := r.Read(); e != nil {
		return scanNumberFail(r, e)
	}

	frac, e := scanInt(r)
	if e != nil {
		return scanNumberFail(r, e)
	}

	return sig + string(dot) + frac, token.Number, nil
}

func scanInt(r Reader) (string, error) {
	val := []rune{}

	for r.More() {
		ru, e := r.Peek()
		if e != nil {
			return scanIntFail(r, e)
		}

		if !isDigit(ru) {
			break
		}

		if _, e = r.Read(); e != nil {
			return scanIntFail(r, e)
		}
		val = append(val, ru)
	}

	return string(val), nil
}

func scanString(r Reader) (string, token.TokenType, error) {

	var str []rune

	ru, e := scanStringDelim(r)
	if e != nil {
		return scanStringFail(r, e)
	}
	str = append(str, ru)

	strBody, e := scanStringBody(r)
	if e != nil {
		return scanStringFail(r, e)
	}

	str = append(str, strBody...)

	ru, e = scanStringDelim(r)
	if e != nil {
		return scanStringFail(r, e)
	}
	str = append(str, ru)

	return string(str), token.String, nil
}

func scanStringDelim(r Reader) (rune, error) {
	ru, e := r.Read()
	if e != nil {
		return scanStringDelimFail(r, nil)
	}

	if ru != StringDelim {
		return rune(0), unterminatedString(r)
	}

	return ru, nil
}

func scanStringBody(r Reader) ([]rune, error) {
	strBody := []rune{}
	escaped := false

	for r.More() {
		ru, e := r.Peek()
		if e != nil {
			return scanStringBodyFail(r, e)
		}

		isDelim := ru == StringDelim
		if !escaped && isDelim {
			break
		}

		_, e = r.Read()
		if e != nil {
			return scanStringBodyFail(r, e)
		}

		escaped = !escaped && ru == StringEscape
		strBody = append(strBody, ru)
	}

	if escaped {
		return nil, unterminatedString(r)
	}

	return strBody, nil
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

func scanNewlineFail(r Reader, e error) (string, token.TokenType, error) {
	return "", token.Unknown, err.Pos(r.Pos(), e, "Failed to scan newline")
}

func scanSpaceFail(r Reader, e error) (string, token.TokenType, error) {
	return "", token.Unknown, err.Pos(r.Pos(), e, "Failed to scan spaces")
}

func scanNumberFail(r Reader, e error) (string, token.TokenType, error) {
	return "", token.Unknown, err.Pos(r.Pos(), e, "Failed to scan number")
}

func scanStringFail(r Reader, e error) (string, token.TokenType, error) {
	return "", token.Unknown, err.Pos(r.Pos(), e, "Failed to scan string")
}

func scanStringDelimFail(r Reader, e error) (rune, error) {
	return rune(0), err.Pos(r.Pos(), e, "Failed to scan string delimiter")
}

func scanStringBodyFail(r Reader, e error) ([]rune, error) {
	return []rune{}, err.Pos(r.Pos(), e, "Failed to scan string body")
}

func scanWordFail(r Reader, e error) (string, token.TokenType, error) {
	return "", token.Unknown, err.Pos(r.Pos(), e, "Failed to scan word")
}

func scanIntFail(r Reader, e error) (string, error) {
	return "", err.Pos(r.Pos(), e, "Failed to scan integer")
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

func unterminatedString(r Reader) error {
	return err.Pos(r.Pos(), nil, "Unterminated string")
}
