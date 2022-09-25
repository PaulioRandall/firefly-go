package rinser

import (
	//"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type tokenStream interface {
	access() token.Token
	accept(token.TokenType) bool
	expect(token.TokenType) error
}

type tkStream struct {
	idx int
	tks []token.Token
}

func newTkStream(tks []token.Token) tkStream {
	return tkStream{
		tks: tks,
	}
}
