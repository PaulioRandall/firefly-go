// Package aligner aligns list and map literals on to a single line
package aligner

import (
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type TokenReader interface {
	More() bool
	Read() token.Token
	Peek() token.Token
}

type tokenList []token.Token

func (tl *tokenList) append(tk token.Token) {
	*tl = append(*tl, tk)
}

func AlignAll(tr TokenReader) []token.Token {
	var tl tokenList

	for tr.More() {
		tk := tr.Read()
		tl.append(tk)

		if closer := getCloserFor(tk.TokenType); closer != token.Unknown {
			alignBlock(tr, &tl, closer)
		}
	}

	return []token.Token(tl)
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

func alignBlock(tr TokenReader, tl *tokenList, closer token.TokenType) {
	for first := true; tr.More(); first = false {
		tk := tr.Read()

		if closer := getCloserFor(tk.TokenType); closer != token.Unknown {
			tl.append(tk)
			alignBlock(tr, tl, closer)
			continue
		}

		if tk.TokenType == closer {
			tl.append(tk)
			return
		}

		if tk.TokenType != token.Newline {
			tl.append(tk)
			continue
		}

		if !first && tr.Peek().TokenType != closer {
			tk.TokenType = token.Comma
			tl.append(tk)
		}
	}
}
