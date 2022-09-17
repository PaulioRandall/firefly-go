package token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, tk Token, isKeyword bool) {
	require.True(t, len(tk.String()) > 0)
	require.Equal(t, isKeyword, tk.IsKeyword())
}

func Test_Unknown(t *testing.T) {
	require.Equal(t, 0, len(Unknown.String()))
	require.False(t, Unknown.IsKeyword())
}

func Test_If(t *testing.T)    { doTest(t, If, true) }
func Test_For(t *testing.T)   { doTest(t, For, true) }
func Test_Watch(t *testing.T) { doTest(t, Watch, true) }
func Test_When(t *testing.T)  { doTest(t, When, true) }
func Test_E(t *testing.T)     { doTest(t, E, true) }
func Test_F(t *testing.T)     { doTest(t, F, true) }
func Test_End(t *testing.T)   { doTest(t, End, true) }
