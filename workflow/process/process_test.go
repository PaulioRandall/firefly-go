package process

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
)

var zero = rune(0)

func when[T comparable](t *testing.T, given []T, p ProcessItem[T]) ([]T, error) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[T]()
	e := Process[T](r, w, p)
	return w.List(), e
}

func Test_1(t *testing.T) {
	given := []rune("a")
	forAcceptProcess := func(prev, curr, next rune) (rune, error) {
		return curr, nil
	}

	act, e := when(t, given, forAcceptProcess)
	exp := []rune("a")

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_2(t *testing.T) {
	given := []rune("abc")
	forAcceptProcess := func(prev, curr, next rune) (rune, error) {
		return curr, nil
	}

	act, e := when(t, given, forAcceptProcess)
	exp := []rune("abc")

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_3(t *testing.T) {
	given := []rune("abc")
	forRejectProcess := func(prev, curr, next rune) (rune, error) {
		return rune(0), nil
	}

	act, e := when(t, given, forRejectProcess)
	exp := []rune(nil)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_4(t *testing.T) {
	given := []rune("abc")
	forModifyProcess := func(prev, curr, next rune) (rune, error) {
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

	act, e := when(t, given, forModifyProcess)
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
