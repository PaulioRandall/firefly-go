package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func Test_1_general(t *testing.T) {
	r := newBR(
		tok(token.String, `""`),
	)

	varMatcher := func(tt token.TokenType) bool {
		return tt == token.Identifier
	}

	isMatch := match(r, varMatcher)

	require.False(t, isMatch)
}

func Test_2_general(t *testing.T) {
	r := newBR(
		tok(token.String, `""`),
	)

	stringMatcher := func(tt token.TokenType) bool {
		return tt == token.String
	}

	isMatch := match(r, stringMatcher)

	require.True(t, isMatch)
}

func Test_3_general(t *testing.T) {
	r := newBR(
		tok(token.String, `""`),
	)

	isMatch := is(r, token.Identifier)

	require.False(t, isMatch)
}

func Test_4_general(t *testing.T) {
	r := newBR(
		tok(token.String, `""`),
	)

	isMatch := is(r, token.String)

	require.True(t, isMatch)
}

func Test_5_general(t *testing.T) {
	r := newBR()
	accepted := accept(r, token.Identifier)

	require.False(t, accepted)
}

func Test_6_general(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	r := newBR(given...)
	accepted := accept(r, token.Number)

	require.False(t, accepted)
	require.True(t, r.More())
}

func Test_7_general(t *testing.T) {
	given := []token.Token{
		tok(token.Identifier, "a"),
	}

	r := newBR(given...)
	accepted := accept(r, token.Identifier)

	require.True(t, accepted)
	require.Equal(t, given[0], r.Prev())
	require.False(t, r.More())
}

func Test_8_general(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "1"),
	}

	r := newBR(given...)
	accept(r, token.String)
	accepted := accept(r, token.Number)

	require.True(t, accepted)
	require.Equal(t, given[1], r.Prev())
	require.False(t, r.More())
}

func Test_9_general(t *testing.T) {
	r := newBR()

	require.Panics(t, func() {
		_ = expect(r, token.EQU)
	})
}

func Test_10_general(t *testing.T) {
	r := newBR()

	defer func() {
		e := recover()
		require.NotNil(t, e)

		isUnexpectedEOF := err.Is(e.(error), UnexpectedEOF)
		require.True(t, isUnexpectedEOF)
	}()

	_ = expect(r, token.EQU)
}

func Test_11_general(t *testing.T) {
	r := newBR(
		tok(token.NEQ, "!="),
	)

	require.Panics(t, func() {
		_ = expect(r, token.EQU)
	})
}

func Test_12_general(t *testing.T) {
	r := newBR(
		tok(token.NEQ, "!="),
	)

	defer func() {
		e := recover()
		require.NotNil(t, e)

		isUnexpectedToken := err.Is(e.(error), UnexpectedToken)
		require.True(t, isUnexpectedToken)
	}()

	_ = expect(r, token.EQU)
}

func Test_13_general(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	r := newBR(given...)
	tk := expect(r, token.String)

	require.Equal(t, given[0], tk)
	require.Equal(t, given[0], r.Prev())
	require.False(t, r.More())
}

func Test_14_general(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "1"),
	}

	r := newBR(given...)
	_ = expect(r, token.String)
	require.True(t, r.More())
}

func Test_15_general(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "1"),
	}

	r := newBR(given...)
	_ = expect(r, token.String)
	tk := expect(r, token.Number)

	require.Equal(t, given[1], tk)
	require.Equal(t, given[1], r.Prev())
	require.False(t, r.More())
}
