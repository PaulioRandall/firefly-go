package scanner

import (
	"github.com/PaulioRandall/firefly-go/pkg/err"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

type tokenBuilder struct {
	r     Reader
	start token.Pos
	tt    token.TokenType
	val   []rune
}

func (tb *tokenBuilder) any() error {
	_, e := tb.acceptFunc(func(have rune) bool {
		return true
	})
	return e
}

func (tb *tokenBuilder) accept(want rune) (bool, error) {
	return tb.acceptFunc(func(have rune) bool {
		return have == want
	})
}

func (tb *tokenBuilder) acceptFunc(f func(rune) bool) (bool, error) {
	if !tb.r.More() {
		return false, nil
	}

	have, e := tb.r.Peek()
	if e != nil {
		return false, e
	}

	if !f(have) {
		return false, nil
	}

	if _, e = tb.r.Read(); e != nil {
		return false, e
	}

	tb.add(have)
	return true, nil
}

func (tb *tokenBuilder) expect(
	want rune,
	errMsg string,
	args ...interface{}) error {

	matcher := func(have rune) bool { return have == want }
	return tb.expectFunc(matcher, errMsg, args...)
}

func (tb *tokenBuilder) expectFunc(
	f func(rune) bool,
	errMsg string,
	args ...interface{}) error {

	found, e := tb.acceptFunc(f)
	if e != nil {
		return err.Pos(tb.r.Pos(), e, "Failed to read from stream")
	}

	if !found {
		return err.Pos(tb.r.Pos(), e, errMsg, args...)
	}

	return nil
}

func (tb *tokenBuilder) add(ru ...rune) {
	tb.val = append(tb.val, ru...)
}

func (tb tokenBuilder) str() string {
	return string(tb.val)
}
