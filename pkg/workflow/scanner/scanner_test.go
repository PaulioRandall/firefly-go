package scanner

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/debug"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func assertToken(t *testing.T, given string, expType token.TokenType) {
	r := inout.NewListReader([]rune(given))
	w := inout.NewListWriter[token.Token]()

	e := Scan(r, w)

	require.Nil(t, e, "Expected %q but got %+v", expType.String(), debug.String(e))
	require.Equal(t, 1, len(w.List()))

	actType := w.List()[0].TokenType
	require.Equal(t, expType, actType,
		"Expected %q but got %q", expType.String(), actType.String(),
	)
}

func assertScan(t *testing.T, given string, exp []token.Token) {
	r := inout.NewListReader([]rune(given))
	w := inout.NewListWriter[token.Token]()

	e := Scan(r, w)
	require.Nil(t, e, "%+v", e)
	tokentest.RequireEqual(t, exp, w.List())
}

func assertError(t *testing.T, given string, exp error) {
	r := inout.NewListReader([]rune(given))
	w := inout.NewListWriter[token.Token]()

	e := Scan(r, w)
	require.True(t, err.Is(e, exp), "Expected %+v", exp.Error())
}

func Test_1(t *testing.T) {
	assertScan(t, "", nil)
}

func Test_2(t *testing.T) {
	given := "\n"

	exp := []token.Token{
		token.MakeTokenAt(
			token.Newline,
			given,
			pos.At(0, 0, 0),
		),
	}

	assertScan(t, given, exp)
}

func Test_3(t *testing.T) {
	assertToken(t, "+", token.Add)
}

func Test_4(t *testing.T) {
	assertToken(t, "-", token.Sub)
}

func Test_5(t *testing.T) {
	assertToken(t, "*", token.Mul)
}

func Test_6(t *testing.T) {
	assertToken(t, "/", token.Div)
}

func Test_7(t *testing.T) {
	assertToken(t, "%", token.Mod)
}

func Test_8(t *testing.T) {
	assertToken(t, "<", token.LT)
}

func Test_9(t *testing.T) {
	assertToken(t, ">", token.GT)
}

func Test_10(t *testing.T) {
	assertToken(t, "<=", token.LTE)
}

func Test_11(t *testing.T) {
	assertToken(t, ">=", token.GTE)
}

func Test_12(t *testing.T) {
	assertToken(t, "==", token.EQU)
}

func Test_13(t *testing.T) {
	assertToken(t, "!=", token.NEQ)
}

func Test_14(t *testing.T) {
	assertToken(t, "=", token.Assign)
}

func Test_15(t *testing.T) {
	assertToken(t, ":=", token.Define)
}

func Test_16(t *testing.T) {
	assertToken(t, ":", token.Colon)
}

func Test_17(t *testing.T) {
	assertToken(t, ";", token.Terminator)
}

func Test_18(t *testing.T) {
	assertToken(t, ",", token.Comma)
}

func Test_19(t *testing.T) {
	assertToken(t, "@", token.Spell)
}

func Test_20(t *testing.T) {
	assertToken(t, "(", token.ParenOpen)
}

func Test_21(t *testing.T) {
	assertToken(t, ")", token.ParenClose)
}

func Test_22(t *testing.T) {
	assertToken(t, "{", token.BraceOpen)
}

func Test_23(t *testing.T) {
	assertToken(t, "}", token.BraceClose)
}

func Test_24(t *testing.T) {
	assertToken(t, "[", token.BracketOpen)
}

func Test_25(t *testing.T) {
	assertToken(t, "]", token.BracketClose)
}

func Test_33(t *testing.T) {
	assertToken(t, `""`, token.String)
}

func Test_34(t *testing.T) {
	assertToken(t, `"a"`, token.String)
}

func Test_35(t *testing.T) {
	assertToken(t, `"abc"`, token.String)
}

func Test_36(t *testing.T) {
	assertToken(t, `"   "`, token.String)
}

func Test_37(t *testing.T) {
	assertToken(t, `"\\"`, token.String)
}

func Test_38(t *testing.T) {
	assertToken(t, `"\\\\\\"`, token.String)
}

func Test_39(t *testing.T) {
	assertToken(t, `"\"\"\""`, token.String)
}

func Test_40(t *testing.T) {
	assertToken(t, " ", token.Space)
}

func Test_41(t *testing.T) {
	assertToken(t, "\t", token.Space)
}

func Test_42(t *testing.T) {
	assertToken(t, "\v", token.Space)
}

func Test_43(t *testing.T) {
	assertToken(t, "\r", token.Space)
}

func Test_44(t *testing.T) {
	assertToken(t, "\f", token.Space)
}

func Test_45(t *testing.T) {
	assertToken(t, "  \t\v \f\r   \v\v\t", token.Space)
}

func Test_50(t *testing.T) {
	assertToken(t, "0", token.Number)
}

func Test_51(t *testing.T) {
	assertToken(t, "0.00000", token.Number)
}

func Test_52(t *testing.T) {
	assertToken(t, "0.1", token.Number)
}

func Test_53(t *testing.T) {
	assertToken(t, "1", token.Number)
}

func Test_54(t *testing.T) {
	assertToken(t, "1.1234567890", token.Number)
}

func Test_55(t *testing.T) {
	assertToken(t, "123456789.987654321", token.Number)
}

func Test_56(t *testing.T) {
	assertToken(t, "9", token.Number)
}

func Test_60(t *testing.T) {
	assertToken(t, "abc", token.Identifier)
}

func Test_61(t *testing.T) {
	assertToken(t, "abc_xyz", token.Identifier)
}

func Test_62(t *testing.T) {
	assertToken(t, "forest", token.Identifier)
}

func Test_63(t *testing.T) {
	assertToken(t, "For", token.Identifier)
}

func Test_64(t *testing.T) {
	assertToken(t, "FOR", token.Identifier)
}

func Test_65(t *testing.T) {
	assertToken(t, "e", token.Identifier)
}

func Test_70(t *testing.T) {
	assertToken(t, "if", token.If)
}

func Test_71(t *testing.T) {
	assertToken(t, "for", token.For)
}

func Test_72(t *testing.T) {
	assertToken(t, "in", token.In)
}

func Test_73(t *testing.T) {
	assertToken(t, "if", token.If)
}

func Test_74(t *testing.T) {
	assertToken(t, "watch", token.Watch)
}

func Test_75(t *testing.T) {
	assertToken(t, "when", token.When)
}

func Test_76(t *testing.T) {
	assertToken(t, "is", token.Is)
}

func Test_77(t *testing.T) {
	assertToken(t, "E", token.Expr)
}

func Test_78(t *testing.T) {
	assertToken(t, "F", token.Func)
}

func Test_79(t *testing.T) {
	assertToken(t, "true", token.True)
}

func Test_80(t *testing.T) {
	assertToken(t, "false", token.False)
}

func Test_81(t *testing.T) {
	assertToken(t, `"\\"`, token.String)
}

func Test_82(t *testing.T) {
	assertToken(t, "//", token.Comment)
}

func Test_83(t *testing.T) {
	assertToken(t, "// abc", token.Comment)
}

func Test_100(t *testing.T) {
	assertError(t, "~", ErrUnknownSymbol)
}

func Test_101(t *testing.T) {
	assertError(t, `"`, ErrUnterminatedString)
}

func Test_102(t *testing.T) {
	assertError(t, `"""`, ErrUnterminatedString)
}

func Test_103(t *testing.T) {
	assertError(t, `"\`, ErrUnterminatedString)
}

func Test_104(t *testing.T) {
	assertError(t, `"\"`, ErrUnterminatedString)
}

func Test_105(t *testing.T) {
	assertError(t, `"\\\"`, ErrUnterminatedString)
}

func Test_106(t *testing.T) {
	assertError(t, "=!", ErrUnknownSymbol)
}

func Test_107(t *testing.T) {
	assertError(t, ".", ErrUnknownSymbol)
}

func Test_108(t *testing.T) {
	assertError(t, "0.", ErrMissingFractional)
}

func Test_109(t *testing.T) {
	assertError(t, "0.a", ErrMissingFractional)
}

func Test_200(t *testing.T) {
	given := "x = 1"

	gen := tokentest.NewTokenGenerator()
	exp := []token.Token{
		gen(token.Identifier, "x"),
		gen(token.Space, " "),
		gen(token.Assign, "="),
		gen(token.Space, " "),
		gen(token.Number, "1"),
	}

	assertScan(t, given, exp)
}

func Test_201(t *testing.T) {
	given := strings.Join([]string{
		`x = true`,
		`y, z = 123.456, "string"`,
		``,
		`// A function`,
		`f := F(a, b) c, d`,
		`	when a`,
		`		is 1: @println("one")`,
		`		a == b: @println("a == b")`,
		`		true: @println("meh")`,
		`	end`,
		`end`,
		``,
	}, "\n")

	gen := tokentest.NewTokenGenerator()
	exp := []token.Token{
		// `x = true`
		gen(token.Identifier, "x"),
		gen(token.Space, " "),
		gen(token.Assign, "="),
		gen(token.Space, " "),
		gen(token.True, "true"),
		gen(token.Newline, "\n"),
		// `y, z = 1, "string"`
		gen(token.Identifier, "y"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.Identifier, "z"),
		gen(token.Space, " "),
		gen(token.Assign, "="),
		gen(token.Space, " "),
		gen(token.Number, "123.456"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.String, `"string"`),
		gen(token.Newline, "\n"),
		// ``
		gen(token.Newline, "\n"),
		// `// A function`
		gen(token.Comment, "// A function"),
		gen(token.Newline, "\n"),
		// `f := F(a, b) c, d {`
		gen(token.Identifier, "f"),
		gen(token.Space, " "),
		gen(token.Define, ":="),
		gen(token.Space, " "),
		gen(token.Func, "F"),
		gen(token.ParenOpen, "("),
		gen(token.Identifier, "a"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.Identifier, "b"),
		gen(token.ParenClose, ")"),
		gen(token.Space, " "),
		gen(token.Identifier, "c"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.Identifier, "d"),
		gen(token.Newline, "\n"),
		// `	when a`
		gen(token.Space, "\t"),
		gen(token.When, "when"),
		gen(token.Space, " "),
		gen(token.Identifier, "a"),
		gen(token.Newline, "\n"),
		// ` 	is 1: @println("one")`
		gen(token.Space, "\t\t"),
		gen(token.Is, "is"),
		gen(token.Space, " "),
		gen(token.Number, "1"),
		gen(token.Colon, ":"),
		gen(token.Space, " "),
		gen(token.Spell, "@"),
		gen(token.Identifier, "println"),
		gen(token.ParenOpen, "("),
		gen(token.String, `"one"`),
		gen(token.ParenClose, ")"),
		gen(token.Newline, "\n"),
		// `		a == b: @println("b")`,
		gen(token.Space, "\t\t"),
		gen(token.Identifier, "a"),
		gen(token.Space, " "),
		gen(token.EQU, "=="),
		gen(token.Space, " "),
		gen(token.Identifier, "b"),
		gen(token.Colon, ":"),
		gen(token.Space, " "),
		gen(token.Spell, "@"),
		gen(token.Identifier, "println"),
		gen(token.ParenOpen, "("),
		gen(token.String, `"a == b"`),
		gen(token.ParenClose, ")"),
		gen(token.Newline, "\n"),
		// `		true: @println("meh")`
		gen(token.Space, "\t\t"),
		gen(token.True, "true"),
		gen(token.Colon, ":"),
		gen(token.Space, " "),
		gen(token.Spell, "@"),
		gen(token.Identifier, "println"),
		gen(token.ParenOpen, "("),
		gen(token.String, `"meh"`),
		gen(token.ParenClose, ")"),
		gen(token.Newline, "\n"),
		// `	}`
		gen(token.Space, "\t"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
		// `}`
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	assertScan(t, given, exp)
}
