package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

func aud(given ...token.Token) *auditor.Auditor {
	return auditor.NewAuditor(inout.NewListReader(given))
}

func Test_1_general(t *testing.T) {
	a := aud(
		tok(token.String, `""`),
	)

	varMatcher := func(tt token.TokenType) bool {
		return tt == token.Identifier
	}

	isMatch := doesNextMatch(a, varMatcher)

	require.False(t, isMatch)
}

func Test_2_general(t *testing.T) {
	a := aud(
		tok(token.String, `""`),
	)

	stringMatcher := func(tt token.TokenType) bool {
		return tt == token.String
	}

	isMatch := doesNextMatch(a, stringMatcher)

	require.True(t, isMatch)
}

func Test_3_general(t *testing.T) {
	a := aud(
		tok(token.String, `""`),
	)

	isMatch := isNext(a, token.Identifier)

	require.False(t, isMatch)
}

func Test_4_general(t *testing.T) {
	a := aud(
		tok(token.String, `""`),
	)

	isMatch := isNext(a, token.String)

	require.True(t, isMatch)
}
