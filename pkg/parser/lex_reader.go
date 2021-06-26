package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

type lexReader struct {
	lr token.LexemeReader
}

func (r lexReader) More() bool {
	return r.lr.More()
}

func (r lexReader) Read() token.Lexeme {
	lx, e := r.lr.Read()
	if e != nil {
		lexemeReadPanic(e)
	}
	return lx
}

func (r lexReader) PutBack(lx token.Lexeme) {
	e := r.lr.PutBack(lx)
	if e != nil {
		lexemeReadPanic(e)
	}
}

func lexemeReadPanic(cause error) {
	parsingPanic(cause, "Lexeme read error")
}
