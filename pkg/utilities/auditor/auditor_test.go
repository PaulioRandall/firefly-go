package auditor

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func tok2(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func aud(given ...token.Token) *Auditor {
	return NewAuditor(inout.NewListReader(given))
}

func Test_1_auditor_accept(t *testing.T) {
	a := aud()

	accepted := a.Accept(token.Identifier)

	require.False(t, accepted)
}

func Test_2_auditor_accept(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	accepted := a.Accept(token.Number)

	require.False(t, accepted)
	require.True(t, a.More())
}

func Test_3_auditor_accept(t *testing.T) {
	a := aud(
		tok2(token.Identifier, "a"),
	)

	accepted := a.Accept(token.Identifier)

	require.True(t, accepted)
	require.Equal(t, tok2(token.Identifier, "a"), a.Prev())
	require.False(t, a.More())
}

func Test_4_auditor_accept(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
		tok2(token.Number, "1"),
	)

	a.Accept(token.String)
	accepted := a.Accept(token.Number)

	require.True(t, accepted)
	require.Equal(t, tok2(token.Number, "1"), a.Prev())
	require.False(t, a.More())
}

func Test_5_auditor_expect(t *testing.T) {
	a := aud()

	require.Panics(t, func() {
		_ = a.Expect(token.EQU)
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

	_ = a.Expect(token.EQU)
}

func Test_7_auditor_expect(t *testing.T) {
	a := aud(
		tok2(token.NEQ, "!="),
	)

	require.Panics(t, func() {
		_ = a.Expect(token.EQU)
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

	_ = a.Expect(token.EQU)
}

func Test_9_auditor_expect(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	tk := a.Expect(token.String)

	require.Equal(t, tok2(token.String, `""`), tk)
	require.Equal(t, tok2(token.String, `""`), a.Prev())
	require.False(t, a.More())
}

func Test_10_auditor_expect(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
		tok2(token.Number, "1"),
	)

	_ = a.Expect(token.String)
	require.True(t, a.More())
}

func Test_11_auditor_expect(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
		tok2(token.Number, "1"),
	)

	_ = a.Expect(token.String)
	tk := a.Expect(token.Number)

	require.Equal(t, tok2(token.Number, "1"), tk)
	require.Equal(t, tok2(token.Number, "1"), a.Prev())
	require.False(t, a.More())
}

func Test_16_auditor_peekNext(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	tk := a.Peek()

	require.Equal(t, tok2(token.String, `""`), tk)
	require.True(t, a.More())
}

func Test_17_auditor_peekNext(t *testing.T) {
	a := aud()

	require.Panics(t, func() {
		_ = a.Peek()
	})
}

func Test_18_auditor_readNext(t *testing.T) {
	a := aud()

	require.Panics(t, func() {
		_ = a.Read()
	})
}

func Test_19_auditor_readNext(t *testing.T) {
	a := aud(
		tok2(token.String, `""`),
	)

	tk := a.Read()

	require.Equal(t, tok2(token.String, `""`), tk)
	require.False(t, a.More())
}
