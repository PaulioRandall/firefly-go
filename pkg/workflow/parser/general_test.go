package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

func aud(given ...token.Token) tokenAuditor {
	return auditor.NewAuditor[token.Token](inout.NewListReader(given))
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

func Test_5_general(t *testing.T) {
	a := aud()
	accepted := accept(a, token.Identifier)

	require.False(t, accepted)
}

func Test_6_general(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	a := aud(given...)
	accepted := accept(a, token.Number)

	require.False(t, accepted)
	require.True(t, a.More())
}

func Test_7_general(t *testing.T) {
	given := []token.Token{
		tok(token.Identifier, "a"),
	}

	a := aud(given...)
	accepted := accept(a, token.Identifier)

	require.True(t, accepted)
	require.Equal(t, given[0], a.Prev())
	require.False(t, a.More())
}

func Test_8_general(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "1"),
	}

	a := aud(given...)
	accept(a, token.String)
	accepted := accept(a, token.Number)

	require.True(t, accepted)
	require.Equal(t, given[1], a.Prev())
	require.False(t, a.More())
}

func Test_9_general(t *testing.T) {
	a := aud()

	require.Panics(t, func() {
		_ = expect(a, token.EQU)
	})
}

func Test_10_general(t *testing.T) {
	a := aud()

	defer func() {
		e := recover()
		require.NotNil(t, e)

		isUnexpectedEOF := err.Is(e.(error), UnexpectedEOF)
		require.True(t, isUnexpectedEOF)
	}()

	_ = expect(a, token.EQU)
}

func Test_11_general(t *testing.T) {
	a := aud(
		tok(token.NEQ, "!="),
	)

	require.Panics(t, func() {
		_ = expect(a, token.EQU)
	})
}

func Test_12_general(t *testing.T) {
	a := aud(
		tok(token.NEQ, "!="),
	)

	defer func() {
		e := recover()
		require.NotNil(t, e)

		isUnexpectedToken := err.Is(e.(error), UnexpectedToken)
		require.True(t, isUnexpectedToken)
	}()

	_ = expect(a, token.EQU)
}

func Test_13_general(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	a := aud(given...)
	tk := expect(a, token.String)

	require.Equal(t, given[0], tk)
	require.Equal(t, given[0], a.Prev())
	require.False(t, a.More())
}

func Test_14_auditor_expect(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "1"),
	}

	a := aud(given...)
	_ = expect(a, token.String)
	require.True(t, a.More())
}

func Test_15_general(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "1"),
	}

	a := aud(given...)
	_ = expect(a, token.String)
	tk := expect(a, token.Number)

	require.Equal(t, given[1], tk)
	require.Equal(t, given[1], a.Prev())
	require.False(t, a.More())
}
