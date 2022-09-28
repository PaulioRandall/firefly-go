package compiler

/*
type tokenList []token.Token

func (tl *tokenList) append(tk token.Token) {
	*tl = append(*tl, tk)
}

func Compile(tr TokenReader) []token.Token {
	var tl tokenList

	for tr.More() {
		tk := tr.Read()
		tl.append(tk)

		if closer := getCloserFor(tk.Type); closer != token.Unknown {
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

		if closer := getCloserFor(tk.Type); closer != token.Unknown {
			tl.append(tk)
			alignBlock(tr, tl, closer)
			continue
		}

		if tk.Type == closer {
			tl.append(tk)
			return
		}

		if tk.Type != token.Newline {
			tl.append(tk)
			continue
		}

		if !first && tr.Peek().Type != closer {
			tk.Type = token.Comma
			tl.append(tk)
		}
	}
}
*/
