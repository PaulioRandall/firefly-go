package process

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
)

var zero = rune(0)

func when[T any](t *testing.T, given []T, p ProcessItem[T]) ([]T, error) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[T]()
	e := Process[T](r, w, p)
	return w.List(), e
}

func Test_1(t *testing.T) {
	given := []rune("a")
	forProcess := func(prev, curr, next rune) (rune, bool, error) {
		return curr, true, nil
	}

	act, e := when(t, given, forProcess)
	exp := []rune("a")

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_2(t *testing.T) {
	given := []rune("abc")
	forProcess := func(prev, curr, next rune) (rune, bool, error) {
		return curr, true, nil
	}

	act, e := when(t, given, forProcess)
	exp := []rune("abc")

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_3(t *testing.T) {
	given := []rune("abc")
	forProcess := func(prev, curr, next rune) (rune, bool, error) {
		return curr, false, nil
	}

	act, e := when(t, given, forProcess)
	exp := []rune(nil)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_4(t *testing.T) {
	given := []rune("abc")
	forProcess := func(prev, curr, next rune) (rune, bool, error) {
		switch curr {
		case 'a':
			curr = 'x'
		case 'b':
			curr = 'y'
		case 'c':
			curr = 'z'
		}
		return curr, true, nil
	}

	act, e := when(t, given, forProcess)
	exp := []rune("xyz")

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_5(t *testing.T) {
	given := []rune("abcd")
	forProcess := func(prev, curr, next rune) (rune, bool, error) {
		switch curr {
		case 'b':
			return 'y', true, nil
		case 'c':
			return rune(0), false, nil
		default:
			return curr, true, nil
		}
	}

	act, e := when(t, given, forProcess)
	exp := []rune("ayd")

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
