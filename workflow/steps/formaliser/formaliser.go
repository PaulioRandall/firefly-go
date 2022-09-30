package formaliser

import (
	"github.com/PaulioRandall/firefly-go/workflow/readers/tokenreader"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type tokenList []token.Token

func (tl *tokenList) append(tk token.Token) {
	*tl = append(*tl, tk)
}

func (tl *tokenList) last() token.Token {
	if i := tl.indexOfLast(); i > -1 {
		return (*tl)[i]
	}
	return token.Token{}
}

func (tl *tokenList) indexOfLast() int {
	return len(*tl) - 1
}

var zeroTk token.Token

func Formalise(tr tokenreader.TokenReader) []token.Token {
	var tl tokenList

	for tr.More() {
		tk := formalise(&tl, tr.Read())

		if tk != zeroTk {
			tl.append(tk)
		}
	}

	return []token.Token(tl)
}

func formalise(tl *tokenList, tk token.Token) token.Token {

	switch {
	case tk.TokenType == token.Newline:
		if preventsNewlineTermination(tl.last()) {
			tk = zeroTk
		} else {
			tk.TokenType = token.Terminator
		}
	}

	return tk
}

func preventsNewlineTermination(tk token.Token) bool {
	switch tk.TokenType {
	case token.Add, token.Sub, token.Mul, token.Div, token.Mod:
		return true
	default:
		return false
	}
}
