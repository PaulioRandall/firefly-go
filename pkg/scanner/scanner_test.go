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

func Test_ScanAll_1(t *testing.T) {
	r := readers.NewRuneStringReader("")

	act, e := ScanAll(r)
	var exp []token.Token

	require.Nil(t, e)
	require.Equal(t, exp, act)
}

func Test_ScanAll_2(t *testing.T) {
	r := readers.NewRuneStringReader("~")
	_, e := ScanAll(r)
	require.NotNil(t, e)
}

func Test_ScanAll_3(t *testing.T) {
	exp := singletonList(token.If, "if", 2)
	doScanTest(t, "if", exp)
}

func Test_ScanAll_4(t *testing.T) {
	exp := singletonList(token.For, "for", 3)
	doScanTest(t, "for", exp)
}

func Test_ScanAll_5(t *testing.T) {
	exp := singletonList(token.Watch, "watch", 5)
	doScanTest(t, "watch", exp)
}

func Test_ScanAll_6(t *testing.T) {
	exp := singletonList(token.When, "when", 4)
	doScanTest(t, "when", exp)
}

func Test_ScanAll_7(t *testing.T) {
	exp := singletonList(token.E, "E", 1)
	doScanTest(t, "E", exp)
}

func Test_ScanAll_8(t *testing.T) {
	exp := singletonList(token.F, "F", 1)
	doScanTest(t, "F", exp)
}

func Test_ScanAll_9(t *testing.T) {
	exp := singletonList(token.End, "end", 3)
	doScanTest(t, "end", exp)
}

func Test_ScanAll_10(t *testing.T) {
	exp := singletonList(token.Var, "abc", 3)
	doScanTest(t, "abc", exp)
}

func Test_ScanAll_11(t *testing.T) {
	exp := singletonList(token.Var, "abc_xyz", 7)
	doScanTest(t, "abc_xyz", exp)
}

func Test_ScanAll_12(t *testing.T) {
	exp := singletonList(token.Var, "_", 1)
	doScanTest(t, "_", exp)
}
