package token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func doTokenTest(t *testing.T, tk Token, isKeyword bool) {
	require.True(t, len(tk.String()) > 0)
	require.Equal(t, isKeyword, tk.IsKeyword())
}

func Test_Token_Unknown(t *testing.T) {
	require.Equal(t, 0, len(Unknown.String()))
	require.False(t, Unknown.IsKeyword())
}

func Test_Token_If(t *testing.T)    { doTokenTest(t, If, true) }
func Test_Token_For(t *testing.T)   { doTokenTest(t, For, true) }
func Test_Token_Watch(t *testing.T) { doTokenTest(t, Watch, true) }
func Test_Token_When(t *testing.T)  { doTokenTest(t, When, true) }
func Test_Token_E(t *testing.T)     { doTokenTest(t, E, true) }
func Test_Token_F(t *testing.T)     { doTokenTest(t, F, true) }
func Test_Token_End(t *testing.T)   { doTokenTest(t, End, true) }
