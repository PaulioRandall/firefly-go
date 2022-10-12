package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/err"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func tok2(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func aud(given ...token.Token) *auditor {
	return newAuditor(inout.NewListReader(given))
}

func Test_1_auditor_accept(t *testing.T) {
	a := aud()

	accepted := a.accept(token.Identifier)

	require.False(t, accepted)
}

func Test_2_auditor_accept(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	accepted := a.accept(token.Number)

	require.False(t, accepted)
	require.True(t, a.more())
}

func Test_3_auditor_accept(t *testing.T) {
	a := aud(
		tok2(token.Identifier, "a"),
	)

	accepted := a.accept(token.Identifier)

	require.True(t, accepted)
	require.Equal(t, tok2(token.Identifier, "a"), a.getPrev())
	require.False(t, a.more())
}

func Test_4_auditor_accept(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
		tok2(token.Number, "1"),
	)

	a.accept(token.String)
	accepted := a.accept(token.Number)

	require.True(t, accepted)
	require.Equal(t, tok2(token.Number, "1"), a.getPrev())
	require.False(t, a.more())
}

func Test_5_auditor_expect(t *testing.T) {
	a := aud()

	require.Panics(t, func() {
		_ = a.expect(token.EQU)
	})
}

func Test_6_auditor_expect(t *testing.T) {
	a := aud()

	defer func() {
		e := recover()
		require.NotNil(t, e)

		isUnexpectedEOF := err.Is(e.(error), UnexpectedEOF)
		require.True(t, isUnexpectedEOF)
	}()

	_ = a.expect(token.EQU)
}

func Test_7_auditor_expect(t *testing.T) {
	a := aud(
		tok2(token.NEQ, "!="),
	)

	require.Panics(t, func() {
		_ = a.expect(token.EQU)
	})
}

func Test_8_auditor_expect(t *testing.T) {
	a := aud(
		tok2(token.NEQ, "!="),
	)

	defer func() {
		e := recover()
		require.NotNil(t, e)

		isUnexpectedToken := err.Is(e.(error), UnexpectedToken)
		require.True(t, isUnexpectedToken)
	}()

	_ = a.expect(token.EQU)
}

func Test_9_auditor_expect(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	tk := a.expect(token.String)

	require.Equal(t, tok2(token.String, `""`), tk)
	require.Equal(t, tok2(token.String, `""`), a.getPrev())
	require.False(t, a.more())
}

func Test_10_auditor_expect(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
		tok2(token.Number, "1"),
	)

	_ = a.expect(token.String)
	require.True(t, a.more())
}

func Test_11_auditor_expect(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
		tok2(token.Number, "1"),
	)

	_ = a.expect(token.String)
	tk := a.expect(token.Number)

	require.Equal(t, tok2(token.Number, "1"), tk)
	require.Equal(t, tok2(token.Number, "1"), a.getPrev())
	require.False(t, a.more())
}

func Test_12_auditor_doesNextMatch(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	varMatcher := func(tt token.TokenType) bool {
		return tt == token.Identifier
	}

	isMatch := a.doesNextMatch(varMatcher)

	require.False(t, isMatch)
}

func Test_13_auditor_doesNextMatch(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	stringMatcher := func(tt token.TokenType) bool {
		return tt == token.String
	}

	isMatch := a.doesNextMatch(stringMatcher)

	require.True(t, isMatch)
}

func Test_14_auditor_isNext(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	isMatch := a.isNext(token.Identifier)

	require.False(t, isMatch)
}

func Test_15_auditor_isNext(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	isMatch := a.isNext(token.String)

	require.True(t, isMatch)
}

func Test_16_auditor_peekNext(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	tk := a.peekNext()

	require.Equal(t, tok2(token.String, `""`), tk)
	require.True(t, a.more())
}

func Test_17_auditor_peekNext(t *testing.T) {
	a := aud()

	require.Panics(t, func() {
		_ = a.peekNext()
	})
}

func Test_18_auditor_readNext(t *testing.T) {
	a := aud()

	require.Panics(t, func() {
		_ = a.readNext()
	})
}

func Test_19_auditor_readNext(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	tk := a.readNext()

	require.Equal(t, tok2(token.String, `""`), tk)
	require.False(t, a.more())
}
