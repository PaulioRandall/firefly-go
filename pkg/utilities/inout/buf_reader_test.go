package inout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func newBR(given ...rune) BufReader[rune] {
	lr := NewListReader[rune](given)
	return NewBufReader[rune](lr)
}

func Test_1_bufReader(t *testing.T) {
	given := []rune{'a'}

	r := newBR(given...)
	act, e := r.Peek()

	require.Nil(t, e, "%+v", e)
	require.Equal(t, given[0], act)
	require.True(t, r.More())
}

func Test_2_bufReader(t *testing.T) {
	r := newBR()
	_, e := r.Peek()
	requireEOF(t, e)
}

func Test_3_bufReader(t *testing.T) {
	r := newBR()
	_, e := r.Read()
	requireEOF(t, e)
}

func Test_4_bufReader(t *testing.T) {
	given := []rune{'a'}

	r := newBR(given...)
	act, e := r.Read()

	require.Nil(t, e, "%+v", e)
	require.Equal(t, given[0], act)
	require.False(t, r.More())
}

func Test_5_bufReader(t *testing.T) {
	r := newBR()
	exp := 'a'

	r.Putback(exp)
	act, _ := r.Peek()

	require.Equal(t, exp, act)
	require.True(t, r.More())
}

func Test_6_bufReader(t *testing.T) {
	r := newBR('a', 'b')
	exp := 'c'

	r.Putback(exp)
	act, _ := r.Peek()

	require.Equal(t, exp, act)
	require.True(t, r.More())
}

func Test_7_bufReader(t *testing.T) {
	r := newBR('a', 'b')
	act := r.Prev()
	require.Equal(t, rune(0), act)
}

func Test_8_bufReader(t *testing.T) {
	r := newBR('a', 'b')

	actRead, _ := r.Read()
	actPrev := r.Prev()

	require.Equal(t, 'a', actRead)
	require.Equal(t, 'a', actPrev)
}
