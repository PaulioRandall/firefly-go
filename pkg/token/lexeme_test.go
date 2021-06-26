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
			lex(TK_NUMBER, "1"),
			lex(TK_ADD, "+"),
			lex(TK_ADD, "+"),
			lex(TK_NUMBER, "2"),
		},
	}

	r := &lexemeReader{
		lxs: []Lexeme{
			lex(TK_NUMBER, "1"),
			lex(TK_ADD, "+"),
			lex(TK_NUMBER, "2"),
		},
	}

	_, e := r.Read()
	require.Nil(t, e, "%+v", e)

	lx, e := r.Read()
	require.Nil(t, e, "%+v", e)

	_ = r.PutBack(lx)

	require.Equal(t, exp, r)
}
