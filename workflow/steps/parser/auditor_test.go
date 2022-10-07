package parser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"

	"github.com/PaulioRandall/firefly-go/workflow/token/tokentest"
)

func audTok(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func aud(given ...token.Token) auditor {
	return auditor{
		TokenReader: inout.NewListReader(given),
	}
}

func Test_1_auditor_accept(t *testing.T) {
	a := aud()

	accepted := a.accept(token.Var)

	require.False(t, accepted)
}

func Test_2_auditor_accept(t *testing.T) {
	a := aud(
		audTok(token.String, `""`),
	)

	accepted := a.accept(token.Number)

	require.False(t, accepted)
	require.True(t, a.More())
}

func Test_3_auditor_accept(t *testing.T) {
	a := aud(
		audTok(token.Var, "a"),
	)

	accepted := a.accept(token.Var)

	require.True(t, accepted)
	require.Equal(t, audTok(token.Var, "a"), a.access())
	require.False(t, a.More())
}

func Test_4_auditor_accept(t *testing.T) {
	a := aud(
		audTok(token.String, `""`),
		audTok(token.Number, "1"),
	)

	a.accept(token.String)
	accepted := a.accept(token.Number)

	require.True(t, accepted)
	require.Equal(t, audTok(token.Number, "1"), a.access())
	require.False(t, a.More())
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

		isUnexpectedEOF := errors.Is(e.(error), err.UnexpectedEOF)
		require.True(t, isUnexpectedEOF)
	}()

	_ = a.expect(token.EQU)
}

func Test_7_auditor_expect(t *testing.T) {
	a := aud(
		audTok(token.NEQ, "!="),
	)

	require.Panics(t, func() {
		_ = a.expect(token.EQU)
	})
}

func Test_8_auditor_expect(t *testing.T) {
	a := aud(
		audTok(token.NEQ, "!="),
	)

	defer func() {
		e := recover()
		require.NotNil(t, e)

		isUnexpectedToken := errors.Is(e.(error), err.UnexpectedToken)
		require.True(t, isUnexpectedToken)
	}()

	_ = a.expect(token.EQU)
}

func Test_9_auditor_expect(t *testing.T) {
	a := aud(
		audTok(token.String, `""`),
	)

	tk := a.expect(token.String)

	require.Equal(t, audTok(token.String, `""`), tk)
	require.Equal(t, audTok(token.String, `""`), a.access())
	require.False(t, a.More())
}

func Test_10_auditor_expect(t *testing.T) {
	a := aud(
		audTok(token.String, `""`),
		audTok(token.Number, "1"),
	)

	_ = a.expect(token.String)
	require.True(t, a.More())
}

func Test_11_auditor_expect(t *testing.T) {
	a := aud(
		audTok(token.String, `""`),
		audTok(token.Number, "1"),
	)

	_ = a.expect(token.String)
	tk := a.expect(token.Number)

	require.Equal(t, audTok(token.Number, "1"), tk)
	require.Equal(t, audTok(token.Number, "1"), a.access())
	require.False(t, a.More())
}
