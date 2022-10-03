package scanner

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var ErrNotFound = errors.New("Symbol not found")

type tokenBuilder struct {
	in    Input
	start token.Pos
	pos   token.Pos
	tt    token.TokenType
	val   []rune
}

func (tb *tokenBuilder) err(
	cause error,
	errMsg string,
	args ...interface{}) error {

	return err.AtPos(tb.pos, cause, errMsg, args...)
}

func (tb tokenBuilder) more() bool {
	return tb.in.More()
}

func (tb *tokenBuilder) peek() (rune, error) {
	return tb.in.Peek()
}

func (tb *tokenBuilder) read() (rune, error) {
	return tb.in.Read()
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
	if !tb.in.More() {
		return false, nil
	}

	have, e := tb.in.Peek()
	if e != nil {
		return false, e
	}

	if !f(have) {
		return false, nil
	}

	if _, e = tb.in.Read(); e != nil {
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

	if !tb.in.More() {
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
	tb.pos.IncString(string(ru))
}

func (tb *tokenBuilder) build() token.Token {

	s := string(tb.val)

	rng := token.MakeRange(tb.start, tb.pos)
	tk := token.MakeToken(tb.tt, s, rng)

	tb.start = tb.pos
	tb.tt = token.Unknown
	tb.val = nil

	return tk
}
