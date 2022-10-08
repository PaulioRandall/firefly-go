package container

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LinkedQueue_enforceTypes(t *testing.T) {
	var _ Queue[rune] = &LinkedQueue[rune]{}
}

func Test_1_LinkedQueue(t *testing.T) {
	require.Panics(t, func() {
		q := LinkedQueue[rune]{}
		q.Next()
	})
}

func Test_2_LinkedQueue(t *testing.T) {
	require.Panics(t, func() {
		q := LinkedQueue[rune]{}
		q.Last()
	})
}

func Test_3_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')

	require.True(t, q.More())
	require.Equal(t, 1, q.Len())
	require.Equal(t, 'a', q.Next())
	require.Equal(t, 'a', q.Last())
}

func Test_4_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')
	q.Add('b')
	q.Add('c')

	require.Equal(t, 3, q.Len())
	require.Equal(t, 'a', q.Next())
	require.Equal(t, 'c', q.Last())
}

func Test_5_LinkedQueue(t *testing.T) {
	require.Panics(t, func() {
		q := LinkedQueue[rune]{}
		q.Take()
	})
}

func Test_6_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')
	act := q.Take()

	require.Equal(t, 'a', act)
	require.Equal(t, 0, q.Len())
	require.False(t, q.More())

	require.Panics(t, func() {
		q.Next()
	})

	require.Panics(t, func() {
		q.Last()
	})
}

func Test_7_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')
	q.Add('b')
	q.Add('c')

	act := q.Take()

	require.Equal(t, 'a', act)
	require.Equal(t, 'b', q.Next())
	require.Equal(t, 'c', q.Last())

	require.True(t, q.More())
	require.Equal(t, 2, q.Len())
}

func Test_8_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')
	q.Add('b')
	q.Add('c')

	_ = q.Take()
	_ = q.Take()
	act := q.Take()

	require.Equal(t, 'c', act)
	require.False(t, q.More())
	require.Equal(t, 0, q.Len())
}

func Test_9_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Push('a')

	require.Equal(t, 'a', q.Next())
	require.Equal(t, 'a', q.Last())

	require.True(t, q.More())
	require.Equal(t, 1, q.Len())
}

func Test_10_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Push('a')
	q.Push('b')
	q.Push('c')

	require.Equal(t, 3, q.Len())
	require.Equal(t, 'c', q.Next())
	require.Equal(t, 'a', q.Last())
}

func Test_11_LinkedQueue(t *testing.T) {
	require.Panics(t, func() {
		q := LinkedQueue[rune]{}
		q.TakeBack()
	})
}

func Test_12_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')
	act := q.TakeBack()

	require.Equal(t, 'a', act)
	require.False(t, q.More())
	require.Equal(t, 0, q.Len())

	require.Panics(t, func() {
		q.Next()
	})

	require.Panics(t, func() {
		q.Last()
	})
}

func assertQueue[T any](t *testing.T, exp []T, q Queue[T]) {

	if len(exp) == 0 {
		require.True(t, q.Empty())
		return
	}

	lastIdx := len(exp) - 1
	require.Equal(t, exp[0], q.Next())
	require.Equal(t, exp[lastIdx], q.Last())

	for _, want := range exp {
		require.True(t, q.More())
		require.Equal(t, want, q.Take())
	}
}

func Test_13_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')
	q.Add('b')
	q.Take()
	q.Push('c')
	q.TakeBack()
	q.Push('a')
	q.TakeBack()
	q.Add('d')

	exp := []rune{
		'a',
		'd',
	}

	assertQueue[rune](t, exp, &q)
}
