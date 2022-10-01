package workflow

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/readers/runereader"
	//"github.com/PaulioRandall/firefly-go/workflow/token"
)

func Test_1_Workflow(t *testing.T) {
	rr := runereader.FromString("")

	act, e := Parse(rr)

	require.True(t, errors.Is(e, err.EOF))
	require.Empty(t, act)
}

/*
func Test_2_Workflow(t *testing.T) {
	rr := runereader.FromString("")

	act, e := Parse(rr)

	require.Nil(t, e, "%+v", e)
	require.Empty(t, act)
}
*/
