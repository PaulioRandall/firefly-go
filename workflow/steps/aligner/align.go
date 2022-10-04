// Package aligner aligns list and map literals on to a single line
package aligner

import (
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type TokenReader interface {
	More() bool
	Read() (token.Token, error)
	Peek() (token.Token, error)
}

type TokenWriter interface {
	Write(token.Token) error
}

func Align(r TokenReader, w TokenWriter) error {

	for r.More() {
		tk, e := r.Read()
		if e != nil {
			return e // TODO: Wrap error using current/previous token position
		}

		if e = writeToken(w, tk, "Failed to write token"); e != nil {
			return e // TODO: Wrap error using current/previous token position
		}

		if closer := getCloserFor(tk.TokenType); closer != token.Unknown {
			if e = alignBlock(r, w, closer); e != nil {
				return e // TODO: Wrap error using current/previous token position
			}
		}
	}

	return nil
}

func getCloserFor(opener token.TokenType) token.TokenType {
	switch opener {
	case token.BracketOpen:
		return token.BracketClose
	case token.BraceOpen:
		return token.BraceClose
	case token.ParenOpen:
		return token.ParenClose
	}

	return token.Unknown
}

func alignBlock(r TokenReader, w TokenWriter, closer token.TokenType) error {
	for first := true; r.More(); first = false {
		tk, e := r.Read()
		if e != nil {
			return e // TODO: Wrap error using current/previous token position
		}

		if innerCloser := getCloserFor(tk.TokenType); innerCloser != token.Unknown {
			if e = nestedBlock(r, w, tk, innerCloser); e != nil {
				return e // TODO: Wrap error using current/previous token position
			}
			continue
		}

		if tk.TokenType == closer {
			return writeToken(w, tk, "Failed to close block")
		}

		if tk.TokenType != token.Newline {
			if e = writeToken(w, tk, "Failed to write token"); e != nil {
				return e
			}
			continue
		}

		if first {
			continue
		}

		next, e := r.Peek()
		if e != nil {
			return e // TODO: Wrap error using current/previous token position
		}

		if next.TokenType != closer {
			tk.TokenType = token.Comma
			e = writeToken(w, tk, "Failed converting token from newline to comma")
			if e != nil {
				return e
			}
		}
	}

	return nil
}

func writeToken(w TokenWriter, tk token.Token, msg string, args ...any) error {
	if e := w.Write(tk); e != nil {
		return e // TODO: Wrap error using current/previous token position
	}
	return nil
}

func nestedBlock(r TokenReader, w TokenWriter, opener token.Token, closer token.TokenType) error {
	if e := writeToken(w, opener, "Failed to block opener"); e != nil {
		return e
	}

	if e := alignBlock(r, w, closer); e != nil {
		return e // TODO: Wrap error using current/previous token position
	}

	return nil
}
