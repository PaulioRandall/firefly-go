package rinser

/*
import (
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func RinseAll(r TokenReader) ([]token.Token, error) {
	var (
		tk  token.Token
		tks []token.Token
		sc  = New(r)
		e   error
	)

	for sc != nil {
		tk, sc, e = sc()

		if e != nil {
			return nil, err.Pos(r.Pos(), e, "Failed to scan all tokens")
		}

		tks = append(tks, tk)
	}

	return tks, nil
}
*/
