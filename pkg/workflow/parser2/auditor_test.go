package parser2

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

func Test_1_auditor(t *testing.T) {
	a := newAud(
		tok(token.String, `""`),
	)

	varMatcher := func(tt token.TokenType) bool {
		return tt == token.Identifier
	}

	isMatch := a.match(varMatcher)

	require.False(t, isMatch)
}

func Test_2_auditor(t *testing.T) {
	a := newAud(
		tok(token.String, `""`),
	)

	stringMatcher := func(tt token.TokenType) bool {
		return tt == token.String
	}

	isMatch := a.match(stringMatcher)

	require.True(t, isMatch)
}

func Test_3_auditor(t *testing.T) {
	a := newAud(
		tok(token.String, `""`),
	)

	isMatch := a.is(token.Identifier)

	require.False(t, isMatch)
}

func Test_4_auditor(t *testing.T) {
	a := newAud(
		tok(token.String, `""`),
	)

	isMatch := a.is(token.String)

	require.True(t, isMatch)
}

func Test_5_auditor(t *testing.T) {
	a := newAud()
	accepted := a.accept(token.Identifier)

	require.False(t, accepted)
}

func Test_6_auditor(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	a := newAud(given...)
	accepted := a.accept(token.Number)

	require.False(t, accepted)
	require.True(t, a.More())
}

func Test_7_auditor(t *testing.T) {
	given := []token.Token{
		tok(token.Identifier, "a"),
	}

	a := newAud(given...)
	accepted := a.accept(token.Identifier)

	require.True(t, accepted)
	require.Equal(t, given[0], a.Prev())
	require.False(t, a.More())
}

func Test_8_auditor(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "1"),
	}

	a := newAud(given...)
	a.accept(token.String)
	accepted := a.accept(token.Number)

	require.True(t, accepted)
	require.Equal(t, given[1], a.Prev())
	require.False(t, a.More())
}

func Test_9_auditor(t *testing.T) {
	a := newAud()

	require.Panics(t, func() {
		_ = a.expect(token.EQU)
	})
}

func Test_10_auditor(t *testing.T) {
	a := newAud()

	defer func() {
		e := recover()
		require.NotNil(t, e)

		isUnexpectedEOF := err.Is(e.(error), ErrUnexpectedEOF)
		require.True(t, isUnexpectedEOF)
	}()

	_ = a.expect(token.EQU)
}

func Test_11_auditor(t *testing.T) {
	a := newAud(
		tok(token.NEQ, "!="),
	)

	require.Panics(t, func() {
		_ = a.expect(token.EQU)
	})
}

func Test_12_auditor(t *testing.T) {
	a := newAud(
		tok(token.NEQ, "!="),
	)

	defer func() {
		e := recover()
		require.NotNil(t, e)

		isErrUnexpectedToken := err.Is(e.(error), ErrUnexpectedToken)
		require.True(t, isErrUnexpectedToken)
	}()

	_ = a.expect(token.EQU)
}

func Test_13_auditor(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	a := newAud(given...)
	tk := a.expect(token.String)

	require.Equal(t, given[0], tk)
	require.Equal(t, given[0], a.Prev())
	require.False(t, a.More())
}

func Test_14_auditor(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "1"),
	}

	a := newAud(given...)
	_ = a.expect(token.String)
	require.True(t, a.More())
}

func Test_15_auditor(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "1"),
	}

	a := newAud(given...)
	_ = a.expect(token.String)
	tk := a.expect(token.Number)

	require.Equal(t, given[1], tk)
	require.Equal(t, given[1], a.Prev())
	require.False(t, a.More())
}

func Test_16_auditor(t *testing.T) {
	a := newAud()
	_, accepted := a.acquire(token.Identifier)

	require.False(t, accepted)
}

func Test_17_auditor(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	a := newAud(given...)
	_, accepted := a.acquire(token.Number)

	require.False(t, accepted)
	require.True(t, a.More())
}

func Test_18_auditor(t *testing.T) {
	given := []token.Token{
		tok(token.Identifier, "a"),
	}

	a := newAud(given...)
	tk, accepted := a.acquire(token.Identifier)

	require.True(t, accepted)
	require.Equal(t, given[0], tk)
	require.Equal(t, tk, a.Prev())
	require.False(t, a.More())
}

func Test_19_auditor(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "1"),
	}

	a := newAud(given...)
	a.accept(token.String)
	tk, accepted := a.acquire(token.Number)

	require.True(t, accepted)
	require.Equal(t, given[1], tk)
	require.Equal(t, tk, a.Prev())
	require.False(t, a.More())
}
