package inout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_1_outputList(t *testing.T) {
	out := ToList[string]()
	exp := []string{"abc"}

	require.Nil(t, out.Write("abc"))
	require.Equal(t, exp, out.List())
}

func Test_2_outputList(t *testing.T) {
	out := ToList[string]()
	exp := []string{"x", "y", "z"}

	require.Nil(t, out.Write("x", "y", "z"))
	require.Equal(t, exp, out.List())
}
