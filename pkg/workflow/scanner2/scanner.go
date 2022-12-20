// Package scanner converts a string of runes into meaningful tokens.
package scanner2

import (
	//"unicode"

	"github.com/PaulioRandall/go-trackerr"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

const (
	StringEscape = '\\'
	StringDelim  = '"'
	Newline      = '\n'
	Terminator   = ';'
)

var (
	ErrScanning           = trackerr.Checkpoint("Token scanning failed")
	ErrUnknownSymbol      = trackerr.Track("Unknown symbol")
	ErrUnterminatedString = trackerr.Track("Unterminated string")
	ErrMissingFractional  = trackerr.Track("Missing fractional part of number")
	zeroToken             token.Token
)

type ReaderOfRunes interface {
	More() bool
	Read() rune
	Putback(rune)
}

type WriterOfTokens interface {
	Write(token.Token) error
}

func Scan(r ReaderOfRunes, w WriterOfTokens) error {
	for r.More() {
		tk, e := scanNext(r)
		if e != nil {
			return ErrScanning.Wrap(e)
		}

		if e := w.Write(tk); e != nil {
			return ErrScanning.CausedBy(e, "Could not write scanned token")
		}
	}

	return nil
}

func scanNext(r ReaderOfRunes) (token.Token, error) {
	zero := token.Token{}
	ru := r.Read()

	switch {
	case isNewline(ru):
		return makeToken(token.Newline, string(ru)), nil
	default:
		r.Putback(ru)
	}

	return zero, ErrUnknownSymbol
}

func makeToken(tt token.TokenType, v string) token.Token {
	return token.Token{
		TokenType: tt,
		Value:     v,
	}
}

func isNewline(ru rune) bool {
	return ru == '\n'
}
