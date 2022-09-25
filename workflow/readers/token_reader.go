package readers

import (
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type tokenReader struct {
	idx int
	tks []token.Token
}

func NewTokenReader(tks []token.Token) tokenReader {
	return tokenReader{
		tks: tks,
	}
}

func (tr tokenReader) More() bool {
	return tr.idx < len(tr.tks)
}

func (tr *tokenReader) Read() token.Token {
	tk := tr.tks[tr.idx]
	tr.idx++
	return tk
}

func (tr *tokenReader) Peek() token.Token {
	return tr.tks[tr.idx]
}
