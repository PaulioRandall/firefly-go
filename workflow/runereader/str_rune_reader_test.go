package runereader

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func Test_1_scrollReader_Peek(t *testing.T) {
	rr := FromString("")
	_, e := rr.Peek()
	require.Equal(t, err.EOF, e)
}

func Test_2_scrollReader_Peek(t *testing.T) {
	rr := FromString("abc")

	ru1, e1 := rr.Peek()
	ru2, e2 := rr.Peek()

	require.Nil(t, e1)
	require.Nil(t, e2)

	require.Equal(t, 'a', ru1)
	require.Equal(t, 'a', ru2)

	require.True(t, rr.More())
	expPos := token.MakePos(0, 0, 0)
	require.Equal(t, expPos, rr.Pos())
}

func Test_1_scrollReader_Read(t *testing.T) {
	rr := FromString("")
	_, e := rr.Read()
	require.Equal(t, err.EOF, e)
}

func Test_2_scrollReader_Read(t *testing.T) {
	rr := FromString("abc")

	expPos := token.MakePos(0, 0, 0)
	require.Equal(t, expPos, rr.Pos())

	ru, e := rr.Read()
	require.Nil(t, e)
	require.Equal(t, 'a', ru)
	require.True(t, rr.More())
	expPos = token.MakePos(1, 0, 1)
	require.Equal(t, expPos, rr.Pos())

	ru, e = rr.Read()
	require.Nil(t, e)
	require.Equal(t, 'b', ru)
	require.True(t, rr.More())
	expPos = token.MakePos(2, 0, 2)
	require.Equal(t, expPos, rr.Pos())

	ru, e = rr.Read()
	require.Nil(t, e)
	require.Equal(t, 'c', ru)
	require.False(t, rr.More())
	expPos = token.MakePos(3, 0, 3)
	require.Equal(t, expPos, rr.Pos())

	_, e = rr.Read()
	require.Equal(t, err.EOF, e)
}

func Test_3_scrollReader_Read(t *testing.T) {
	rr := FromString("a\nx")

	rr.Read() // a

	ru, e := rr.Read() // \n
	require.Nil(t, e)
	require.Equal(t, '\n', ru)
	require.True(t, rr.More())
	expPos := token.MakePos(2, 1, 0)
	require.Equal(t, expPos, rr.Pos())

	ru, e = rr.Read() // x
	require.Nil(t, e)
	require.Equal(t, 'x', ru)
	require.False(t, rr.More())
	expPos = token.MakePos(3, 1, 1)
	require.Equal(t, expPos, rr.Pos())

	_, e = rr.Read()
	require.Equal(t, err.EOF, e)
}
