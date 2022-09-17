package readers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_More_1(t *testing.T) {
	r := FromString("")
	require.False(t, r.More())

	r = FromString("abc")
	require.True(t, r.More())
}

func Test_Peek_1(t *testing.T) {
	r := FromString("")

	ru, e := r.Peek()

	require.Nil(t, e)
	require.Equal(t, rune(0), ru)
}

func Test_Peek_2(t *testing.T) {
	r := FromString("abc")

	ru1, e := r.Peek()
	require.Nil(t, e)

	ru2, e := r.Peek()
	require.Nil(t, e)

	require.Equal(t, 'a', ru1)
	require.Equal(t, 'a', ru2)
}

func Test_Read_1(t *testing.T) {
	r := FromString("")
	_, e := r.Read()
	require.NotNil(t, e)
}

func Test_Read_2(t *testing.T) {
	r := FromString("abc")

	ru1, e := r.Read()
	require.Nil(t, e)

	ru2, e := r.Read()
	require.Nil(t, e)

	require.Equal(t, 'a', ru1)
	require.Equal(t, 'b', ru2)
}
