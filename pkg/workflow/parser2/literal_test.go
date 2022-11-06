package parser2

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/debug"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func doLiteralTest(t *testing.T, tt token.TokenType, val string, exp any) {
	defer func() {
		if e := recover(); e != nil {
			require.Fail(t, debug.String(e))
		}
	}()

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(tt, val),
	}

	expect := ast.Literal{
		Value: exp,
	}

	r := inout.NewListReader(given)
	br := inout.NewBufReader[token.Token](r)
	actual := parseLiteral(br)

	require.Equal(t, expect, actual, debug.String(actual))
}

func Test_parseLiteral_1(t *testing.T) {
	doLiteralTest(t, token.Bool, "true", true)
}

func Test_parseLiteral_2(t *testing.T) {
	doLiteralTest(t, token.Bool, "false", false)
}

func Test_parseLiteral_3(t *testing.T) {
	doLiteralTest(t, token.Number, "1", float64(1))
}

func Test_parseLiteral_4(t *testing.T) {
	doLiteralTest(t, token.String, `"abc"`, `abc`)
}
