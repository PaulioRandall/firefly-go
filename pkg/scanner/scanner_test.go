package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/readers"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func singletonList(tt token.TokenType, val string, valLen int) []token.Token {
	return []token.Token{
		token.MakeToken(
			tt,
			val,
			token.MakeInlineRange(0, 0, 0, valLen),
		),
	}
}

func doScanTest(t *testing.T, given string, exp []token.Token) {
	r := readers.NewRuneStringReader(given)

	act, e := ScanAll(r)
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_ScanAll_001(t *testing.T) {
	r := readers.NewRuneStringReader("")

	act, e := ScanAll(r)
	var exp []token.Token

	require.Nil(t, e)
	require.Equal(t, exp, act)
}

func Test_ScanAll_002(t *testing.T) {
	r := readers.NewRuneStringReader("~")
	_, e := ScanAll(r)
	require.NotNil(t, e)
}

func Test_ScanAll_101(t *testing.T) {
	exp := singletonList(token.If, "if", 2)
	doScanTest(t, "if", exp)
}

func Test_ScanAll_102(t *testing.T) {
	exp := singletonList(token.For, "for", 3)
	doScanTest(t, "for", exp)
}

func Test_ScanAll_103(t *testing.T) {
	exp := singletonList(token.Watch, "watch", 5)
	doScanTest(t, "watch", exp)
}

func Test_ScanAll_104(t *testing.T) {
	exp := singletonList(token.When, "when", 4)
	doScanTest(t, "when", exp)
}

func Test_ScanAll_105(t *testing.T) {
	exp := singletonList(token.E, "E", 1)
	doScanTest(t, "E", exp)
}

func Test_ScanAll_106(t *testing.T) {
	exp := singletonList(token.F, "F", 1)
	doScanTest(t, "F", exp)
}

func Test_ScanAll_107(t *testing.T) {
	exp := singletonList(token.End, "end", 3)
	doScanTest(t, "end", exp)
}

func Test_ScanAll_108(t *testing.T) {
	exp := singletonList(token.Var, "abc", 3)
	doScanTest(t, "abc", exp)
}

func Test_ScanAll_109(t *testing.T) {
	exp := singletonList(token.Var, "abc_xyz", 7)
	doScanTest(t, "abc_xyz", exp)
}

func Test_ScanAll_110(t *testing.T) {
	exp := singletonList(token.Var, "_", 1)
	doScanTest(t, "_", exp)
}

func Test_ScanAll_111(t *testing.T) {
	exp := singletonList(token.Var, "forest", 6)
	doScanTest(t, "forest", exp)
}
