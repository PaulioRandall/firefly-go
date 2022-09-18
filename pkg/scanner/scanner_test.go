package scanner

/*
import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/readers"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func Test_ScanAll_1(t *testing.T) {
	r := readers.NewStringRuneReader("")

	act, e := ScanAll(r)
	var exp []token.Lex

	require.Nil(t, e)
	require.Equal(t, exp, act)
}

/*
func Test_ScanAll_2(t *testing.T) {
	r := readers.NewStringRuneReader("if")

	act, e := ScanAll(r)
	exp := []token.Lex{
		token.Lex{
			Token: token.If,
			Value: "if",
			Span:  token.MakeSpan(0, 2, 0, 0),
		},
	}

	require.Nil(t, e)
	require.Equal(t, exp, act)
}
*/
