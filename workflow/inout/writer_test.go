package inout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_enforceTypes_listWriter(t *testing.T) {
	_ = Writer[rune](&listWriter[rune]{})
}

func Test_1_listWriter_Write(t *testing.T) {
	w := NewListWriter[string]()
	exp := []string{"abc"}

	require.Nil(t, w.Write("abc"))
	require.Equal(t, exp, w.List())
}

func Test_2_listWriter_WriteMany(t *testing.T) {
	w := NewListWriter[string]()
	exp := []string{"x", "y", "z"}

	require.Nil(t, w.WriteMany("x", "y", "z"))
	require.Equal(t, exp, w.List())
}
