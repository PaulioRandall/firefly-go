package spells

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrWrongNumberOfParams = err.Trackable("Wrong number of spell parameters")
	ErrWrongParamType      = err.Trackable("Incorrect paramter type")
)

func results(vals ...any) []any {
	return vals
}

func Spell_len(mem Memory, params []any) []any {
	// TODO: Write tests

	if len(params) != 1 {
		panic(ErrWrongNumberOfParams.Trackf(
			"Want 1 argument but got %d",
			len(params),
		))
	}

	p := params[0]

	switch v := p.(type) {
	case string:
		return results(float64(len(v)))
	default:
		panic(ErrWrongParamType.Track("Want type with a length"))
	}
}
