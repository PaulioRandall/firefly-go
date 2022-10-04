package inout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_1_listReader_Peek(t *testing.T) {
	lr := NewListReader([]rune(""))
	_, e := lr.Peek()
	require.Equal(t, EOF, e)
}

func Test_2_listReader_Peek(t *testing.T) {
	lr := NewListReader([]rune("abc"))

	v, e := lr.Peek()
	require.Nil(t, e)
	require.Equal(t, 'a', v)
	require.True(t, lr.More())

	v, e = lr.Peek()
	require.Nil(t, e)
	require.Equal(t, 'a', v)
	require.True(t, lr.More())
}

func Test_1_listReader_Read(t *testing.T) {
	lr := NewListReader([]rune(""))
	_, e := lr.Read()
	require.Equal(t, EOF, e)
}

func Test_2_listReader_Read(t *testing.T) {
	lr := NewListReader([]rune("abc"))

	v, e := lr.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'a', v)
	require.True(t, lr.More())

	v, e = lr.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'b', v)
	require.True(t, lr.More())

	v, e = lr.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'c', v)
	require.False(t, lr.More())

	_, e = lr.Read()
	require.Equal(t, EOF, e)
}
