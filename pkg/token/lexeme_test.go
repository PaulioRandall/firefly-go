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

	exp := &lexemeReader{
		idx: 2,
		lxs: []Lexeme{
			lex(TokenNumber, "1"),
			lex(TokenAdd, "+"),
			lex(TokenAdd, "+"),
			lex(TokenNumber, "2"),
		},
	}

	r := &lexemeReader{
		lxs: []Lexeme{
			lex(TokenNumber, "1"),
			lex(TokenAdd, "+"),
			lex(TokenNumber, "2"),
		},
	}

	_, e := r.Read()
	require.Nil(t, e, "%+v", e)

	lx, e := r.Read()
	require.Nil(t, e, "%+v", e)

	_ = r.PutBack(lx)

	require.Equal(t, exp, r)
}
