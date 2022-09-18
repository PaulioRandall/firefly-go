package token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func doTypeTest(t *testing.T, tt TokenType, isKeyword bool) {
	require.True(t, len(tt.String()) > 0)
	require.Equal(t, isKeyword, tt.IsKeyword())
}

func Test_TokenType_Unknown(t *testing.T) {
	require.Equal(t, 0, len(Unknown.String()))
	require.False(t, Unknown.IsKeyword())
}

func Test_TokenType_If(t *testing.T)    { doTypeTest(t, If, true) }
func Test_TokenType_For(t *testing.T)   { doTypeTest(t, For, true) }
func Test_TokenType_Watch(t *testing.T) { doTypeTest(t, Watch, true) }
func Test_TokenType_When(t *testing.T)  { doTypeTest(t, When, true) }
func Test_TokenType_E(t *testing.T)     { doTypeTest(t, E, true) }
func Test_TokenType_F(t *testing.T)     { doTypeTest(t, F, true) }
func Test_TokenType_End(t *testing.T)   { doTypeTest(t, End, true) }
