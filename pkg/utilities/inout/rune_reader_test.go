package inout

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
)

func newRR(given string) ReaderOfRunes {
	lr := NewListReader([]rune(given))
	return NewReaderOfRunes(lr)
}

func Test_1_runeReader_Peek(t *testing.T) {
	r := newRR("")

	_, e := r.Peek()
	requireError(t, EOF, e, "Expected EOF error")
	require.Empty(t, r.Where())
}

func Test_2_runeReader_Peek(t *testing.T) {
	r := newRR("abc")

	v, e := r.Peek()
	require.Nil(t, e)
	require.Equal(t, 'a', v)
	require.True(t, r.More())
	require.Empty(t, r.Where())

	v, e = r.Peek()
	require.Nil(t, e)
	require.Equal(t, 'a', v)
	require.True(t, r.More())
	require.Empty(t, r.Where())
}

func Test_3_runeReader_Read(t *testing.T) {
	r := newRR("")

	_, e := r.Read()

	requireError(t, EOF, e, "Expected EOF error")
	require.Empty(t, r.Where())
}

func Test_4_runeReader_Read(t *testing.T) {
	r := newRR("ab")

	v, e := r.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'a', v)

	v, e = r.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'b', v)
}

func Test_5_runeReader_Read(t *testing.T) {
	r := newRR("ab")

	_, _ = r.Read()
	require.True(t, r.More())

	_, _ = r.Read()
	require.False(t, r.More())
}

func Test_6_runeReader_Read(t *testing.T) {
	r := newRR("ab")

	_, _ = r.Read()
	require.Equal(t, pos.At(1, 0, 1), r.Where())

	_, _ = r.Read()
	require.Equal(t, pos.At(2, 0, 2), r.Where())
}

func Test_7_runeReader_Read(t *testing.T) {
	lr := NewListReader([]rune("\n"))
	r := NewReaderOfRunes(lr)

	_, _ = r.Read()
	require.Equal(t, pos.At(1, 1, 0), r.Where())
}

func Test_8_runeReader_Read(t *testing.T) {
	r := newRR("\na\nb\n")

	_, _ = r.Read()
	require.Equal(t, pos.At(1, 1, 0), r.Where())

	_, _ = r.Read()
	require.Equal(t, pos.At(2, 1, 1), r.Where())

	_, _ = r.Read()
	require.Equal(t, pos.At(3, 2, 0), r.Where())

	_, _ = r.Read()
	require.Equal(t, pos.At(4, 2, 1), r.Where())

	_, _ = r.Read()
	require.Equal(t, pos.At(5, 3, 0), r.Where())
}
