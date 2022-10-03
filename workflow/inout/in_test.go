package inout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_1_inputList_Peek(t *testing.T) {
	in := FromList([]rune(""))
	_, e := in.Peek()
	require.Equal(t, EOF, e)
}

func Test_2_inputList_Peek(t *testing.T) {
	in := FromList([]rune("abc"))

	v, e := in.Peek()
	require.Nil(t, e)
	require.Equal(t, 'a', v)
	require.True(t, in.More())

	v, e = in.Peek()
	require.Nil(t, e)
	require.Equal(t, 'a', v)
	require.True(t, in.More())
}

func Test_1_inputList_Read(t *testing.T) {
	in := FromList([]rune(""))
	_, e := in.Read()
	require.Equal(t, EOF, e)
}

func Test_2_inputList_Read(t *testing.T) {
	in := FromList([]rune("abc"))

	v, e := in.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'a', v)
	require.True(t, in.More())

	v, e = in.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'b', v)
	require.True(t, in.More())

	v, e = in.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'c', v)
	require.False(t, in.More())

	_, e = in.Read()
	require.Equal(t, EOF, e)
}
