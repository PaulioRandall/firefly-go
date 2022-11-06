package parser2

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/debug"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
	//"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func doExpressionTest(t *testing.T, given []token.Token, exp ast.Node) {
	defer func() {
		if e := recover(); e != nil {
			require.Fail(t, debug.String(e))
		}
	}()

	r := inout.NewListReader(given)
	br := inout.NewBufReader[token.Token](r)
	act := parseExpression(br)

	require.Equal(t, exp, act, debug.String(act))
}
