package container

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LinkedStack_enforceTypes(t *testing.T) {
	var _ Stack[rune] = &LinkedStack[rune]{}
}

func Test_1_LinkedStack(t *testing.T) {
	require.Panics(t, func() {
		st := LinkedStack[rune]{}
		st.Top()
	})
}

func Test_2_LinkedStack(t *testing.T) {
	require.Panics(t, func() {
		st := LinkedStack[rune]{}
		st.Pop()
	})
}

func Test_3_LinkedStack(t *testing.T) {
	st := LinkedStack[rune]{}
	st.Push('a')
	require.Equal(t, 'a', st.Top())
}

func Test_4_LinkedStack(t *testing.T) {
	st := LinkedStack[rune]{}
	st.Push('a')
	st.Push('b')
	require.Equal(t, 'b', st.Top())
}

func Test_5_LinkedStack(t *testing.T) {
	st := LinkedStack[rune]{}

	st.Push('a')
	st.Push('b')

	require.Equal(t, 'b', st.Pop())
	require.Equal(t, 'a', st.Pop())
}
