package scanner

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/pos"
	"github.com/PaulioRandall/firefly-go/workflow/token"

	"github.com/PaulioRandall/firefly-go/workflow/token/tokentest"
)

func assertToken(t *testing.T, given string, expType token.TokenType) {
	in := inout.FromList([]rune(given))
	out := inout.ToList[token.Token]()

	e := Scan(&in, &out)

	require.Nil(t, e, "Expected %q but got %+v", expType.String(), err.Debug(e))
	require.Equal(t, 1, len(out.List()))

	actType := out.List()[0].TokenType
	require.Equal(t, expType, actType,
		"Expected %q but got %q", expType.String(), actType.String(),
	)
}

func assertScan(t *testing.T, given string, exp []token.Token) {
	in := inout.FromList([]rune(given))
	out := inout.ToList[token.Token]()

	e := Scan(&in, &out)
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, out.List())
}

func assertError(t *testing.T, given string, exp error) {
	in := inout.FromList([]rune(given))
	out := inout.ToList[token.Token]()

	e := Scan(&in, &out)
	require.True(t, errors.Is(e, exp), "Expected %+v", exp.Error())
}

func Test_1_Scan(t *testing.T) {
	assertScan(t, "", nil)
}

func Test_7_Scan(t *testing.T) {
	given := "\n"
	exp := []token.Token{
		token.MakeToken(
			token.Newline,
			given,
			pos.MakeRange(
				pos.MakePos(0, 0, 0),
				pos.MakePos(1, 1, 0),
			),
		),
	}

	assertScan(t, given, exp)
}

func Test_10_Scan(t *testing.T) {
	assertToken(t, "=", token.Assign)
}

func Test_11_Scan(t *testing.T) {
	assertToken(t, ":=", token.Define)
}

func Test_12_Scan(t *testing.T) {
	assertToken(t, ";", token.Terminator)
}

func Test_13_Scan(t *testing.T) {
	assertToken(t, ",", token.Comma)
}

func Test_14_Scan(t *testing.T) {
	assertToken(t, ":", token.Colon)
}

func Test_15_Scan(t *testing.T) {
	assertToken(t, "@", token.Spell)
}

func Test_16_Scan(t *testing.T) {
	assertToken(t, "+", token.Add)
}

func Test_17_Scan(t *testing.T) {
	assertToken(t, "-", token.Sub)
}

func Test_18_Scan(t *testing.T) {
	assertToken(t, "*", token.Mul)
}

func Test_19_Scan(t *testing.T) {
	assertToken(t, "/", token.Div)
}

func Test_20_Scan(t *testing.T) {
	assertToken(t, "%", token.Mod)
}

func Test_21_Scan(t *testing.T) {
	assertToken(t, "<", token.LT)
}

func Test_22_Scan(t *testing.T) {
	assertToken(t, ">", token.GT)
}

func Test_23_Scan(t *testing.T) {
	assertToken(t, "<=", token.LTE)
}

func Test_24_Scan(t *testing.T) {
	assertToken(t, ">=", token.GTE)
}

func Test_25_Scan(t *testing.T) {
	assertToken(t, "==", token.EQU)
}

func Test_26_Scan(t *testing.T) {
	assertToken(t, "!=", token.NEQ)
}

func Test_27_Scan(t *testing.T) {
	assertToken(t, "(", token.ParenOpen)
}

func Test_28_Scan(t *testing.T) {
	assertToken(t, ")", token.ParenClose)
}

func Test_29_Scan(t *testing.T) {
	assertToken(t, "{", token.BraceOpen)
}

func Test_30_Scan(t *testing.T) {
	assertToken(t, "}", token.BraceClose)
}

func Test_31_Scan(t *testing.T) {
	assertToken(t, "[", token.BracketOpen)
}

func Test_32_Scan(t *testing.T) {
	assertToken(t, "]", token.BracketClose)
}

func Test_33_Scan(t *testing.T) {
	assertToken(t, `""`, token.String)
}

func Test_34_Scan(t *testing.T) {
	assertToken(t, `"a"`, token.String)
}

func Test_35_Scan(t *testing.T) {
	assertToken(t, `"abc"`, token.String)
}

func Test_36_Scan(t *testing.T) {
	assertToken(t, `"   "`, token.String)
}

func Test_37_Scan(t *testing.T) {
	assertToken(t, `"\\"`, token.String)
}

func Test_38_Scan(t *testing.T) {
	assertToken(t, `"\\\\\\"`, token.String)
}

func Test_39_Scan(t *testing.T) {
	assertToken(t, `"\"\"\""`, token.String)
}

func Test_40_Scan(t *testing.T) {
	assertToken(t, " ", token.Space)
}

func Test_41_Scan(t *testing.T) {
	assertToken(t, "\t", token.Space)
}

func Test_42_Scan(t *testing.T) {
	assertToken(t, "\v", token.Space)
}

func Test_43_Scan(t *testing.T) {
	assertToken(t, "\r", token.Space)
}

func Test_44_Scan(t *testing.T) {
	assertToken(t, "\f", token.Space)
}

func Test_45_Scan(t *testing.T) {
	assertToken(t, "  \t\v \f\r   \v\v\t", token.Space)
}

func Test_50_Scan(t *testing.T) {
	assertToken(t, "0", token.Number)
}

func Test_51_Scan(t *testing.T) {
	assertToken(t, "0.00000", token.Number)
}

func Test_52_Scan(t *testing.T) {
	assertToken(t, "0.1", token.Number)
}

func Test_53_Scan(t *testing.T) {
	assertToken(t, "1", token.Number)
}

func Test_54_Scan(t *testing.T) {
	assertToken(t, "1.1234567890", token.Number)
}

func Test_55_Scan(t *testing.T) {
	assertToken(t, "123456789.987654321", token.Number)
}

func Test_56_Scan(t *testing.T) {
	assertToken(t, "9", token.Number)
}

func Test_60_Scan(t *testing.T) {
	assertToken(t, "abc", token.Var)
}

func Test_61_Scan(t *testing.T) {
	assertToken(t, "abc_xyz", token.Var)
}

func Test_62_Scan(t *testing.T) {
	assertToken(t, "forest", token.Var)
}

func Test_63_Scan(t *testing.T) {
	assertToken(t, "For", token.Var)
}

func Test_64_Scan(t *testing.T) {
	assertToken(t, "FOR", token.Var)
}

func Test_65_Scan(t *testing.T) {
	assertToken(t, "e", token.Var)
}

func Test_70_Scan(t *testing.T) {
	assertToken(t, "if", token.If)
}

func Test_71_Scan(t *testing.T) {
	assertToken(t, "for", token.For)
}

func Test_72_Scan(t *testing.T) {
	assertToken(t, "in", token.In)
}

func Test_73_Scan(t *testing.T) {
	assertToken(t, "if", token.If)
}

func Test_74_Scan(t *testing.T) {
	assertToken(t, "watch", token.Watch)
}

func Test_75_Scan(t *testing.T) {
	assertToken(t, "when", token.When)
}

func Test_76_Scan(t *testing.T) {
	assertToken(t, "is", token.Is)
}

func Test_77_Scan(t *testing.T) {
	assertToken(t, "E", token.E)
}

func Test_78_Scan(t *testing.T) {
	assertToken(t, "F", token.F)
}

func Test_79_Scan(t *testing.T) {
	assertToken(t, "true", token.True)
}

func Test_80_Scan(t *testing.T) {
	assertToken(t, "false", token.False)
}

func Test_81_Scan(t *testing.T) {
	assertToken(t, `"\\"`, token.String)
}

func Test_82_Scan(t *testing.T) {
	assertToken(t, "//", token.Comment)
}

func Test_83_Scan(t *testing.T) {
	assertToken(t, "// abc", token.Comment)
}

func Test_100_Scan(t *testing.T) {
	assertError(t, "~", ErrUnknownSymbol)
}

func Test_101_Scan(t *testing.T) {
	assertError(t, `"`, ErrUnterminatedString)
}

func Test_102_Scan(t *testing.T) {
	assertError(t, `"""`, ErrUnterminatedString)
}

func Test_103_Scan(t *testing.T) {
	assertError(t, `"\`, ErrUnterminatedString)
}

func Test_104_Scan(t *testing.T) {
	assertError(t, `"\"`, ErrUnterminatedString)
}

func Test_105_Scan(t *testing.T) {
	assertError(t, `"\\\"`, ErrUnterminatedString)
}

func Test_106_Scan(t *testing.T) {
	assertError(t, "=!", ErrUnknownSymbol)
}

func Test_107_Scan(t *testing.T) {
	assertError(t, ".", ErrUnknownSymbol)
}

func Test_108_Scan(t *testing.T) {
	assertError(t, "0.", ErrMissingFractional)
}

func Test_109_Scan(t *testing.T) {
	assertError(t, "0.a", ErrMissingFractional)
}

func Test_200_Scan(t *testing.T) {
	given := "x = 1"

	gen := tokentest.NewTokenGenerator()
	exp := []token.Token{
		gen(token.Var, "x"),
		gen(token.Space, " "),
		gen(token.Assign, "="),
		gen(token.Space, " "),
		gen(token.Number, "1"),
	}

	assertScan(t, given, exp)
}

func Test_201_Scan(t *testing.T) {
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
		gen(token.Var, "x"),
		gen(token.Space, " "),
		gen(token.Assign, "="),
		gen(token.Space, " "),
		gen(token.True, "true"),
		gen(token.Newline, "\n"),
		// `y, z = 1, "string"`
		gen(token.Var, "y"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.Var, "z"),
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
		gen(token.Var, "f"),
		gen(token.Space, " "),
		gen(token.Define, ":="),
		gen(token.Space, " "),
		gen(token.F, "F"),
		gen(token.ParenOpen, "("),
		gen(token.Var, "a"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.Var, "b"),
		gen(token.ParenClose, ")"),
		gen(token.Space, " "),
		gen(token.Var, "c"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.Var, "d"),
		gen(token.Newline, "\n"),
		// `	when a`
		gen(token.Space, "\t"),
		gen(token.When, "when"),
		gen(token.Space, " "),
		gen(token.Var, "a"),
		gen(token.Newline, "\n"),
		// ` 	is 1: @println("one")`
		gen(token.Space, "\t\t"),
		gen(token.Is, "is"),
		gen(token.Space, " "),
		gen(token.Number, "1"),
		gen(token.Colon, ":"),
		gen(token.Space, " "),
		gen(token.Spell, "@"),
		gen(token.Var, "println"),
		gen(token.ParenOpen, "("),
		gen(token.String, `"one"`),
		gen(token.ParenClose, ")"),
		gen(token.Newline, "\n"),
		// `		a == b: @println("b")`,
		gen(token.Space, "\t\t"),
		gen(token.Var, "a"),
		gen(token.Space, " "),
		gen(token.EQU, "=="),
		gen(token.Space, " "),
		gen(token.Var, "b"),
		gen(token.Colon, ":"),
		gen(token.Space, " "),
		gen(token.Spell, "@"),
		gen(token.Var, "println"),
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
		gen(token.Var, "println"),
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
