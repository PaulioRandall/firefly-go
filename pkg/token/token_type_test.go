package token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_1_IdentifyWordType(t *testing.T) {
	words := map[string]TokenType{
		"if":    If,
		"for":   For,
		"watch": Watch,
		"when":  When,
		"E":     E,
		"F":     F,
		"end":   End,
		"true":  True,
		"false": False,
		"abc":   Var,
		"For":   Var,
		"e":     Var,
	}

	for val, tt := range words {
		require.Equal(t, tt, IdentifyWordType(val))
	}
}

func Test_2_IdentifyOperatorType(t *testing.T) {
	words := map[string]TokenType{
		"+":  Add,
		"-":  Sub,
		"*":  Mul,
		"/":  Div,
		"%":  Mod,
		"<":  LT,
		">":  GT,
		"<=": LTE,
		">=": GTE,
		"==": EQU,
		"!=": NEQ,
		"~":  Unknown,
		"=>": Unknown,
	}

	for val, tt := range words {
		require.Equal(t, tt, IdentifyOperatorType(val))
	}
}
