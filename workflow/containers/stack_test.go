package containers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_enforce_types(t *testing.T) {
	var _ Stack[rune] = &LinkedStack[rune]{}
}

func Test_1(t *testing.T) {
	require.Panics(t, func() {
		st := LinkedStack[rune]{}
		st.Top()
	})
}

func Test_2(t *testing.T) {
	require.Panics(t, func() {
		st := LinkedStack[rune]{}
		st.Pop()
	})
}

func Test_3(t *testing.T) {
	st := LinkedStack[rune]{}
	st.Push('a')
	require.Equal(t, 'a', st.Top())
}

func Test_4(t *testing.T) {
	st := LinkedStack[rune]{}
	st.Push('a')
	st.Push('b')
	require.Equal(t, 'b', st.Top())
}

func Test_5(t *testing.T) {
	st := LinkedStack[rune]{}

	st.Push('a')
	st.Push('b')

	require.Equal(t, 'b', st.Pop())
	require.Equal(t, 'a', st.Pop())
}
