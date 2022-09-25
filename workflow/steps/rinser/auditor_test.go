package rinser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/readers/tokenreader"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func tok(tt token.TokenType) token.Token {
	return token.MakeToken(tt, "", token.Range{})
}

func newAuditorForTest(tks ...token.Token) auditor {
	return auditor{
		TokenReader: tokenreader.FromList(tks...),
	}
}

func Test_1_auditor_accept(t *testing.T) {
	a := newAuditorForTest()

	act := a.accept(token.Var)

	require.False(t, act)
}

func Test_2_auditor_accept(t *testing.T) {
	a := newAuditorForTest(
		tok(token.String),
	)

	act := a.accept(token.Number)

	require.False(t, act)
	require.True(t, a.More())
}

func Test_3_auditor_accept(t *testing.T) {
	a := newAuditorForTest(
		tok(token.Var),
	)

	act := a.accept(token.Var)

	require.True(t, act)
	require.Equal(t, tok(token.Var), a.access())
	require.False(t, a.More())
}

func Test_4_auditor_accept(t *testing.T) {
	a := newAuditorForTest(
		tok(token.String),
		tok(token.Number),
	)

	a.accept(token.String)
	act := a.accept(token.Number)

	require.True(t, act)
	require.Equal(t, tok(token.Number), a.access())
	require.False(t, a.More())
}

func Test_5_auditor_expect(t *testing.T) {
	a := newAuditorForTest()

	e := a.expect(token.Var)

	require.True(t, errors.Is(e, EOF))
}

func Test_6_auditor_expect(t *testing.T) {
	a := newAuditorForTest(
		tok(token.String),
	)

	e := a.expect(token.Number)

	require.True(t, errors.Is(e, UnexpectedToken))
}

func Test_7_auditor_expect(t *testing.T) {
	a := newAuditorForTest(
		tok(token.String),
	)

	e := a.expect(token.String)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, tok(token.String), a.access())
	require.False(t, a.More())
}

func Test_8_auditor_expect(t *testing.T) {
	a := newAuditorForTest(
		tok(token.String),
		tok(token.Number),
	)

	e := a.expect(token.String)
	require.Nil(t, e, "%+v", e)

	e = a.expect(token.Number)
	require.Nil(t, e, "%+v", e)

	require.Equal(t, tok(token.Number), a.access())
	require.False(t, a.More())
}
