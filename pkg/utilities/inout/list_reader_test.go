package inout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func newLR(given string) Reader[rune] {
	return NewListReader([]rune(given))
}

func Test_1_listReader(t *testing.T) {
	r := newLR("")
	_, e := r.Peek()
	require.Equal(t, EOF, e)
}

func Test_2_listReader(t *testing.T) {
	r := newLR("abc")

	v, e := r.Peek()
	require.Nil(t, e)
	require.Equal(t, 'a', v)
	require.True(t, r.More())

	v, e = r.Peek()
	require.Nil(t, e)
	require.Equal(t, 'a', v)
	require.True(t, r.More())
}

func Test_3_listReader(t *testing.T) {
	r := newLR("")
	_, e := r.Read()
	require.Equal(t, EOF, e)
}

func Test_4_listReader(t *testing.T) {
	r := newLR("abc")

	v, e := r.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'a', v)
	require.True(t, r.More())

	v, e = r.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'b', v)
	require.True(t, r.More())

	v, e = r.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'c', v)
	require.False(t, r.More())

	_, e = r.Read()
	require.Equal(t, EOF, e)
}
