package token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func lex(tk Token, v string) Lexeme {
	return Lexeme{
		Token: tk,
		Value: v,
	}
}

func TestPutBack_1(t *testing.T) {

	exp := &sliceLexemeReader{
		idx: 2,
		lxs: []Lexeme{
			lex(TokenNumber, "1"),
			lex(TokenAdd, "+"),
			lex(TokenAdd, "+"),
			lex(TokenNumber, "2"),
		},
	}

	lr := &sliceLexemeReader{
		lxs: []Lexeme{
			lex(TokenNumber, "1"),
			lex(TokenAdd, "+"),
			lex(TokenNumber, "2"),
		},
	}

	_, e := lr.Read()
	require.Nil(t, e, "%+v", e)

	lx, e := lr.Read()
	require.Nil(t, e, "%+v", e)

	_ = lr.PutBack(lx)

	require.Equal(t, exp, lr)
}
