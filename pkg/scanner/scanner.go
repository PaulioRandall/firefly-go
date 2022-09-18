package scanner

import (
	"fmt"
	"unicode"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

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
			return nil, fmt.Errorf("Failed to scan all tokens: %w", e)
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
		tk, e := scan(r)

		if e != nil {
			return token.Token{}, nil, e
		}

		return tk, NewScanFunc(r), nil
	}
}

func scan(r Reader) (token.Token, error) {

	var (
		zero  token.Token
		tt    token.TokenType
		value string
		start = r.Pos()
	)

	ru, e := r.Peek()
	if e != nil {
		return zero, fmt.Errorf("Failed to scan token: %w", e)
	}

	switch {
	case isWordLetter(ru):
		value, tt, e = scanWord(r)
	default:
		return zero, fmt.Errorf("Unknown symbol %q", ru)
	}

	if e != nil {
		return zero, fmt.Errorf("Failed to scan token: %w", e)
	}

	rng := token.MakeRange(start, r.Pos())
	tk := token.MakeToken(tt, value, rng)
	return tk, nil
}

func scanWord(r Reader) (string, token.TokenType, error) {

	var runes []rune
	err := func(e error) (string, token.TokenType, error) {
		return "", token.Unknown, fmt.Errorf("Failed to scan word: %w", e)
	}

	for r.More() {
		ru, e := r.Peek()
		if e != nil {
			return err(e)
		}

		if !isWordLetter(ru) {
			break
		}

		_, e = r.Read()
		if e != nil {
			return err(e)
		}

		runes = append(runes, ru)
	}

	word := string(runes)
	return word, token.IdentifyWordType(word), nil
}

func isWordLetter(ru rune) bool {
	return unicode.IsLetter(ru) || ru == '_'
}
