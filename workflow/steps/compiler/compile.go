package compiler

import (
	"github.com/PaulioRandall/firefly-go/workflow/readers/tokenreader"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type tokenList []token.Token

func (tl *tokenList) append(tk token.Token) {
	*tl = append(*tl, tk)
}

func Compile(tr tokenreader.TokenReader) []token.Token {
	var tl tokenList

	for tr.More() {
		_ = tr.Read()
	}

	return []token.Token(tl)
}
