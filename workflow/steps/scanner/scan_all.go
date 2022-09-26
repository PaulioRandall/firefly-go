package scanner

import (
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func ScanAll(rr RuneReader) ([]token.Token, error) {
	var (
		tk  token.Token
		tks []token.Token
		sc  = New(rr)
		e   error
	)

	for sc != nil {
		tk, sc, e = sc()

		if e != nil {
			return nil, err.AtPos(rr.Pos(), e, "Failed to scan all tokens")
		}

		tks = append(tks, tk)
	}

	return tks, nil
}
