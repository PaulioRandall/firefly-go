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

func (r lexReader) Peek() token.Lexeme {
	lx, e := r.lr.Peek()
	if e != nil {
		lexemeReadPanic(e)
	}
	return lx
}

func (r lexReader) Read() token.Lexeme {
	lx, e := r.lr.Read()
	if e != nil {
		lexemeReadPanic(e)
	}
	return lx
}

func lexemeReadPanic(cause error) {
	parsingPanic(cause, "Lexeme read error")
}
