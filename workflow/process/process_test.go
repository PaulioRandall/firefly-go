package process

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
)

var zero = rune(0)

func when[In, Out comparable](
	t *testing.T,
	given []In,
	p ProcessItem[In, Out],
) ([]Out, error) {

	r := inout.NewListReader(given)
	w := inout.NewListWriter[Out]()
	e := Process[In, Out](r, w, p)

	return w.List(), e
}

func Test_1(t *testing.T) {
	given := []rune("a")
	acceptEverything := func(prev, curr, next rune) (rune, error) {
		return curr, nil
	}

	act, e := when(t, given, acceptEverything)
	exp := []rune("a")

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_2(t *testing.T) {
	given := []rune("abc")
	acceptEverything := func(prev, curr, next rune) (rune, error) {
		return curr, nil
	}

	act, e := when(t, given, acceptEverything)
	exp := []rune("abc")

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_3(t *testing.T) {
	given := []rune("abc")
	rejectEverything := func(prev, curr, next rune) (rune, error) {
		return rune(0), nil
	}

	act, e := when(t, given, rejectEverything)
	exp := []rune(nil)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_4(t *testing.T) {
	given := []rune("abc")
	mapToXYZ := func(prev, curr, next rune) (rune, error) {
		switch curr {
		case 'a':
			curr = 'x'
		case 'b':
			curr = 'y'
		case 'c':
			curr = 'z'
		}
		return curr, nil
	}

	act, e := when(t, given, mapToXYZ)
	exp := []rune("xyz")

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_5(t *testing.T) {
	given := []rune("abcd")
	forMixedProcess := func(prev, curr, next rune) (rune, error) {
		switch curr {
		case 'b':
			return 'y', nil
		case 'c':
			return rune(0), nil
		default:
			return curr, nil
		}
	}

	act, e := when(t, given, forMixedProcess)
	exp := []rune("ayd")

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_6(t *testing.T) {
	given := []rune("abc")
	invertCaseAndConvToString := func(prev, curr, next rune) (string, error) {
		return string(curr ^ ' '), nil // Invert case
	}

	act, e := when(t, given, invertCaseAndConvToString)
	exp := []string{"A", "B", "C"}

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
