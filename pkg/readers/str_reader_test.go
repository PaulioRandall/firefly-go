package readers

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func Test_Peek_1(t *testing.T) {
	r := NewStringRuneReader("")
	_, e := r.Peek()
	require.NotNil(t, e)
}

func Test_Peek_2(t *testing.T) {
	r := NewStringRuneReader("abc")

	ru, e := r.Peek()
	require.Nil(t, e)
	require.Equal(t, 'a', ru)
	require.True(t, r.More())
	expPos := token.MakePos(0, 0, 0)
	require.Equal(t, expPos, r.Pos())
}

func Test_Peek_3(t *testing.T) {
	r := NewStringRuneReader("abc")

	ru1, e1 := r.Peek()
	ru2, e2 := r.Peek()

	require.Nil(t, e1)
	require.Nil(t, e2)

	require.Equal(t, 'a', ru1)
	require.Equal(t, 'a', ru2)

	require.True(t, r.More())
	expPos := token.MakePos(0, 0, 0)
	require.Equal(t, expPos, r.Pos())
}

func Test_Read_1(t *testing.T) {
	r := NewStringRuneReader("")
	_, e := r.Read()
	require.NotNil(t, e)
}

func Test_Read_2(t *testing.T) {
	r := NewStringRuneReader("abc")

	expPos := token.MakePos(0, 0, 0)
	require.Equal(t, expPos, r.Pos())

	ru, e := r.Read()
	require.Nil(t, e)
	require.Equal(t, 'a', ru)
	require.True(t, r.More())
	expPos = token.MakePos(1, 0, 1)
	require.Equal(t, expPos, r.Pos())

	ru, e = r.Read()
	require.Nil(t, e)
	require.Equal(t, 'b', ru)
	require.True(t, r.More())
	expPos = token.MakePos(2, 0, 2)
	require.Equal(t, expPos, r.Pos())

	ru, e = r.Read()
	require.Nil(t, e)
	require.Equal(t, 'c', ru)
	require.False(t, r.More())
	expPos = token.MakePos(3, 0, 3)
	require.Equal(t, expPos, r.Pos())
}

func Test_Read_3(t *testing.T) {
	r := NewStringRuneReader("a\nx")

	r.Read() // a

	ru, e := r.Read() // \n
	require.Nil(t, e)
	require.Equal(t, '\n', ru)
	require.True(t, r.More())
	expPos := token.MakePos(2, 1, 0)
	require.Equal(t, expPos, r.Pos())

	ru, e = r.Read() // x
	require.Nil(t, e)
	require.Equal(t, 'x', ru)
	require.False(t, r.More())
	expPos = token.MakePos(3, 1, 1)
	require.Equal(t, expPos, r.Pos())
}
