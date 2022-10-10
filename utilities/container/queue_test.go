package container

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LinkedQueue_enforceTypes(t *testing.T) {
	var _ Queue[rune] = &LinkedQueue[rune]{}
}

func assertFirst[T any](t *testing.T, exp T, q Queue[T]) {
	v, ok := q.First()
	require.True(t, ok)
	require.Equal(t, exp, v)
}

func assertLast[T any](t *testing.T, exp T, q Queue[T]) {
	v, ok := q.Last()
	require.True(t, ok)
	require.Equal(t, exp, v)
}

func assertNoFirst[T any](t *testing.T, q Queue[T]) {
	_, ok := q.First()
	require.False(t, ok)
}

func assertNoLast[T any](t *testing.T, q Queue[T]) {
	_, ok := q.Last()
	require.False(t, ok)
}

func Test_1_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}
	_, ok := q.First()
	require.False(t, ok)
}

func Test_2_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}
	_, ok := q.Last()
	require.False(t, ok)
}

func Test_3_LinkedQueue(t *testing.T) {
	q := &LinkedQueue[rune]{}

	q.Add('a')

	require.True(t, q.More())
	require.Equal(t, 1, q.Len())

	assertFirst[rune](t, 'a', q)
	assertLast[rune](t, 'a', q)
}

func Test_4_LinkedQueue(t *testing.T) {
	q := &LinkedQueue[rune]{}

	q.Add('a')
	q.Add('b')
	q.Add('c')

	require.Equal(t, 3, q.Len())

	assertFirst[rune](t, 'a', q)
	assertLast[rune](t, 'c', q)
}

func Test_5_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}
	_, ok := q.Take()
	require.False(t, ok)
}

func Test_6_LinkedQueue(t *testing.T) {
	q := &LinkedQueue[rune]{}

	q.Add('a')
	act, ok := q.Take()

	require.True(t, ok)
	require.Equal(t, 'a', act)
	require.Equal(t, 0, q.Len())
	require.False(t, q.More())

	assertNoFirst[rune](t, q)
	assertNoLast[rune](t, q)
}

func Test_7_LinkedQueue(t *testing.T) {
	q := &LinkedQueue[rune]{}

	q.Add('a')
	q.Add('b')
	q.Add('c')

	act, ok := q.Take()

	require.True(t, ok)
	require.Equal(t, 'a', act)

	assertFirst[rune](t, 'b', q)
	assertLast[rune](t, 'c', q)

	require.True(t, q.More())
	require.Equal(t, 2, q.Len())
}

func Test_8_LinkedQueue(t *testing.T) {
	q := &LinkedQueue[rune]{}

	q.Add('a')
	q.Add('b')
	q.Add('c')

	_, _ = q.Take()
	_, _ = q.Take()
	act, ok := q.Take()

	require.True(t, ok)
	require.Equal(t, 'c', act)
	require.False(t, q.More())
	require.Equal(t, 0, q.Len())
}

func Test_9_LinkedQueue(t *testing.T) {
	q := &LinkedQueue[rune]{}

	q.Return('a')

	assertFirst[rune](t, 'a', q)
	assertLast[rune](t, 'a', q)

	require.True(t, q.More())
	require.Equal(t, 1, q.Len())
}

func Test_10_LinkedQueue(t *testing.T) {
	q := &LinkedQueue[rune]{}

	q.Return('a')
	q.Return('b')
	q.Return('c')

	require.Equal(t, 3, q.Len())

	assertFirst[rune](t, 'c', q)
	assertLast[rune](t, 'a', q)
}

func Test_11_LinkedQueue(t *testing.T) {
	q := LinkedQueue[rune]{}
	_, ok := q.Reclaim()
	require.False(t, ok)
}

func Test_12_LinkedQueue(t *testing.T) {
	q := &LinkedQueue[rune]{}

	q.Add('a')
	act, ok := q.Reclaim()

	require.True(t, ok)
	require.Equal(t, 'a', act)
	require.False(t, q.More())
	require.Equal(t, 0, q.Len())

	assertNoFirst[rune](t, q)
	assertNoLast[rune](t, q)
}

func assertQueue[T any](t *testing.T, exp []T, q Queue[T]) {

	if len(exp) == 0 {
		require.True(t, q.Empty())
		return
	}

	lastIdx := len(exp) - 1
	assertFirst[T](t, exp[0], q)
	assertLast[T](t, exp[lastIdx], q)

	for _, want := range exp {
		act, ok := q.Take()
		require.True(t, ok)
		require.Equal(t, want, act)
	}
}

func Test_13_LinkedQueue(t *testing.T) {
	q := &LinkedQueue[rune]{}

	q.Add('a')
	q.Add('b')
	q.Take()
	q.Return('c')
	q.Reclaim()
	q.Return('a')
	q.Reclaim()
	q.Add('d')

	exp := []rune{
		'a',
		'd',
	}

	assertQueue[rune](t, exp, q)
}
