package container

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LinkedStack_enforceTypes(t *testing.T) {
	var _ Stack[rune] = &LinkedStack[rune]{}
}

func Test_1_LinkedStack(t *testing.T) {
	st := LinkedStack[rune]{}
	_, ok := st.Top()
	require.False(t, ok)
}

func Test_2_LinkedStack(t *testing.T) {
	st := LinkedStack[rune]{}
	_, ok := st.Pop()
	require.False(t, ok)
}

func Test_3_LinkedStack(t *testing.T) {
	st := LinkedStack[rune]{}
	st.Push('a')

	v, ok := st.Top()
	require.True(t, ok)
	require.Equal(t, 'a', v)
}

func Test_4_LinkedStack(t *testing.T) {
	st := LinkedStack[rune]{}
	st.Push('a')
	st.Push('b')

	v, ok := st.Top()
	require.True(t, ok)
	require.Equal(t, 'b', v)
}

func Test_5_LinkedStack(t *testing.T) {
	st := LinkedStack[rune]{}

	st.Push('a')
	st.Push('b')

	v, ok := st.Pop()
	require.True(t, ok)
	require.Equal(t, 'b', v)

	v, ok = st.Pop()
	require.True(t, ok)
	require.Equal(t, 'a', v)
}
