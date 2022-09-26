package rinser

import (
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func RinseAll(tr TokenReader) ([]token.Token, error) {
	var (
		prev, tk  token.Token
		tks       []token.Token
		e         error
		rinseNext = New(tr)
	)

	for {
		prev = tk
		tk, e = rinseNext()

		if e == err.EOF {
			break
		}

		if e != nil {
			return nil, err.AfterToken(prev, e, "Failed to rinse all tokens")
		}

		tks = append(tks, tk)
	}

	return tks, err.EOF
}
