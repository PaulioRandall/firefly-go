package scanner

import (
	"github.com/PaulioRandall/firefly-go/pkg/err"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

type sidekick struct {
	start token.Pos
	tt    token.TokenType
	val   []rune
}

func (sk *sidekick) acceptFunc(r Reader, f func(rune) bool) (bool, error) {
	if !r.More() {
		return false, nil
	}

	have, e := r.Peek()
	if e != nil {
		return false, e
	}

	if !f(have) {
		return false, nil
	}

	if _, e = r.Read(); e != nil {
		return false, e
	}

	sk.add(have)
	return true, nil
}

func (sk *sidekick) accept(r Reader, want rune) (bool, error) {
	return sk.acceptFunc(r, func(have rune) bool {
		return have == want
	})
}

func (sk *sidekick) expectFunc(
	r Reader,
	f func(rune) bool,
	errMsg string,
	args ...interface{}) error {

	found, e := sk.acceptFunc(r, f)
	if e != nil {
		return err.Pos(r.Pos(), e, "Failed to read from stream")
	}

	if !found {
		return err.Pos(r.Pos(), e, errMsg, args...)
	}

	return nil
}

func (sk *sidekick) expect(
	r Reader,
	want rune,
	errMsg string,
	args ...interface{}) error {

	matcher := func(have rune) bool { return have == want }
	return sk.expectFunc(r, matcher, errMsg, args...)
}

func (sk *sidekick) add(ru ...rune) {
	sk.val = append(sk.val, ru...)
}

func (sk sidekick) str() string {
	return string(sk.val)
}
