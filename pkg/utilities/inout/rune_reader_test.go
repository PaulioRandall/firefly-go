package inout

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func Test_enforceTypes_runeReader(t *testing.T) {
	_ = Reader[rune](&readerOfRunes{})
	_ = ReaderOfRunes(&readerOfRunes{})
}

func requireEOF(t *testing.T, e error) {
	require.True(t, err.Is(e, EOF), "Expected EOF error")
}

func Test_1_runeReader_Peek(t *testing.T) {
	lr := NewListReader([]rune(""))
	r := NewReaderOfRunes(lr)

	_, e := r.Peek()
	requireEOF(t, e)
	require.Empty(t, r.Where())
}

func Test_2_runeReader_Peek(t *testing.T) {
	lr := NewListReader([]rune("abc"))
	r := NewReaderOfRunes(lr)

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
	lr := NewListReader([]rune(""))
	r := NewReaderOfRunes(lr)

	_, e := r.Read()

	requireEOF(t, e)
	require.Empty(t, r.Where())
}

func Test_4_runeReader_Read(t *testing.T) {
	lr := NewListReader([]rune("ab"))
	r := NewReaderOfRunes(lr)

	v, e := r.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'a', v)

	v, e = r.Read()
	require.Nil(t, e, "%+v", e)
	require.Equal(t, 'b', v)
}

func Test_5_runeReader_Read(t *testing.T) {
	lr := NewListReader([]rune("ab"))
	r := NewReaderOfRunes(lr)

	_, _ = r.Read()
	require.True(t, r.More())

	_, _ = r.Read()
	require.False(t, r.More())
}

func Test_6_runeReader_Read(t *testing.T) {
	lr := NewListReader([]rune("ab"))
	r := NewReaderOfRunes(lr)

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
	lr := NewListReader([]rune("\na\nb\n"))
	r := NewReaderOfRunes(lr)

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
