package token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Unknown(t *testing.T) {
	require.Equal(t, 0, len(Unknown.String()))
	require.False(t, Unknown.IsKeyword())
}

func Test_If(t *testing.T) {
	require.True(t, len(If.String()) > 0)
	require.True(t, If.IsKeyword())
}
