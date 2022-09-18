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

type ScanFunc func() (tk token.Lex, f ScanFunc, e error)

func NewScanFunc(r Reader) ScanFunc {
	return newScanFunc(r)
}

func ScanAll(r Reader) ([]token.Lex, error) {
	var (
		lx  token.Lex
		lxs []token.Lex
		sc  = newScanFunc(r)
		e   error
	)

	for sc != nil {
		lx, sc, e = sc()

		if e != nil {
			// TODO: wrap error
			return nil, e
		}

		lxs = append(lxs, lx)
	}

	return lxs, nil
}

func newScanFunc(r Reader) ScanFunc {
	if !r.More() {
		return nil
	}

	return func() (token.Lex, ScanFunc, error) {
		lx, e := scan(r)

		if e != nil {
			// TODO: wrap error
			return token.Lex{}, nil, e
		}

		return lx, newScanFunc(r), nil
	}
}

func scan(r Reader) (token.Lex, error) {

	zero := token.Lex{}
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

func scanWord(r Reader) (token.Lex, error) {

	zero := token.Lex{}
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
