package scanner

/*
import (
	"unicode"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

type Reader interface {
	More() bool
	Peek() (rune, error)
	Read() (rune, error)
}

type ScanFunc func() (tk token.Token, f ScanFunc, e error)

func NewScanFunc(r Reader) ScanFunc {
	return newScanFunc(r)
}

func ScanAll(r Reader) ([]token.Token, error) {
	var (
		tk  token.Token
		tks []token.Token
		sc  = newScanFunc(r)
		e   error
	)

	for sc != nil {
		tk, sc, e = sc()

		if e != nil {
			// TODO: wrap error
			return nil, e
		}

		tks = append(tks, tk)
	}

	return tks, nil
}

func newScanFunc(r Reader) ScanFunc {
	if !r.More() {
		return nil
	}

	return func() (token.Token, ScanFunc, error) {
		tk, e := scan(r)

		if e != nil {
			// TODO: wrap error
			return token.Token{}, nil, e
		}

		return tk, newScanFunc(r), nil
	}
}

func scan(r Reader) (token.Token, error) {

	zero := token.Token{}
	ru, e := r.Peek()

	if e != nil {
		// TODO: wrap error
		return zero, e
	}

	switch {
	case isWordLetter(ru):
		return scanWord(r)
	default:
		// TODO: error if no match
		return zero, nil
	}
}

func scanWord(r Reader) (token.Token, error) {

	zero := token.Token{}
	var word []rune

	for r.More() {
		ru, e := r.Peek()

		if e != nil {
			// TODO: wrap error
			return zero, e
		}

		if !isWordLetter(ru) {
			break
		}

		word = append(word, ru)
	}

	return zero, nil
}

func isWordLetter(ru rune) bool {
	return unicode.IsLetter(ru) || ru == '_'
}
*/
