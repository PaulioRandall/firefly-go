package rinser

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func RinseAll(tr TokenReader) []token.Token {
	var (
		tks       []token.Token
		rinseNext = New(tr)
	)

	for {
		tk, e := rinseNext()

		if errors.Is(e, err.EOF) {
			break
		}

		tks = append(tks, tk)
	}

	return tks
}
