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
		q.Front()
	})
}

func Test_2_LinkedQueue(t *testing.T) {
	require.Panics(t, func() {
		q := LinkedQueue[rune]{}
		q.Back()
	})
}

func Test_3_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')

	require.True(t, q.More())
	require.Equal(t, 'a', q.Front())
	require.Equal(t, 'a', q.Back())
}

func Test_4_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')
	q.Add('b')
	q.Add('c')

	require.Equal(t, 'a', q.Front())
	require.Equal(t, 'c', q.Back())
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
	require.False(t, q.More())

	require.Panics(t, func() {
		q.Front()
	})

	require.Panics(t, func() {
		q.Back()
	})
}

func Test_7_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')
	q.Add('b')
	q.Add('c')

	act := q.Take()

	require.Equal(t, 'a', act)
	require.True(t, q.More())
	require.Equal(t, 'b', q.Front())
	require.Equal(t, 'c', q.Back())
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
}

func Test_9_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Push('a')

	require.Equal(t, 'a', q.Front())
	require.Equal(t, 'a', q.Back())
}

func Test_10_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Push('a')
	q.Push('b')
	q.Push('c')

	require.Equal(t, 'c', q.Front())
	require.Equal(t, 'a', q.Back())
}

func Test_11_LinkedQueue(t *testing.T) {
	require.Panics(t, func() {
		q := LinkedQueue[rune]{}
		q.Pop()
	})
}

func Test_12_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}

	q.Add('a')
	act := q.Pop()

	require.Equal(t, 'a', act)
	require.False(t, q.More())

	require.Panics(t, func() {
		q.Front()
	})

	require.Panics(t, func() {
		q.Back()
	})
}

func assertQueue[T any](t *testing.T, exp []T, q Queue[T]) {

	if len(exp) == 0 {
		require.True(t, q.Empty())
		return
	}

	lastIdx := len(exp) - 1
	require.Equal(t, exp[0], q.Front())
	require.Equal(t, exp[lastIdx], q.Back())

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
	q.Pop()
	q.Push('a')
	q.Pop()
	q.Add('d')

	exp := []rune{
		'a',
		'd',
	}

	assertQueue[rune](t, exp, &q)
}
