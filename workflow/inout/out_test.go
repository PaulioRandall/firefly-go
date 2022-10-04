package inout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_1_outputList_Write(t *testing.T) {
	out := ToList[string]()
	exp := []string{"abc"}

	require.Nil(t, out.Write("abc"))
	require.Equal(t, exp, out.List())
}

func Test_2_outputList_WriteMany(t *testing.T) {
	out := ToList[string]()
	exp := []string{"x", "y", "z"}

	require.Nil(t, out.WriteMany("x", "y", "z"))
	require.Equal(t, exp, out.List())
}
