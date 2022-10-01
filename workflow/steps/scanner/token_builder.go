package scanner

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/runereader"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var ErrNotFound = errors.New("Symbol not found")

type tokenBuilder struct {
	rr    runereader.RuneReader
	start token.Pos
	tt    token.TokenType
	val   []rune
}

func (tb *tokenBuilder) err(
	cause error,
	errMsg string,
	args ...interface{}) error {

	return err.AtPos(tb.rr.Pos(), cause, errMsg, args...)
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
	if !tb.rr.More() {
		return false, nil
	}

	have, e := tb.rr.Peek()
	if e != nil {
		return false, e
	}

	if !f(have) {
		return false, nil
	}

	if _, e = tb.rr.Read(); e != nil {
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

	if !tb.rr.More() {
		return err.EOF
	}

	found, e := tb.acceptFunc(f)
	if e != nil {
		return tb.err(e, errMsg, args...)
	}

	if !found {
		return tb.err(ErrNotFound, errMsg, args...)
	}

	return nil
}

func (tb *tokenBuilder) add(ru ...rune) {
	tb.val = append(tb.val, ru...)
}

func (tb tokenBuilder) build() string {
	return string(tb.val)
}
