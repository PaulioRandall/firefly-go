package scanner

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/pos"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var ErrNotFound = errors.New("Unknown symbol")

type runeOutput interface {
	WriteMany(...rune) error
	List() []rune
}

type tokenBuilder struct {
	FFReader
	start pos.Pos
	pos   pos.Pos
	tt    token.TokenType
	out   runeOutput
}

func newTokenBuilder(r FFReader) tokenBuilder {
	return tokenBuilder{
		FFReader: r,
		out:      inout.NewListOutput[rune](),
	}
}

func (tb *tokenBuilder) err(
	cause error,
	errMsg string,
	args ...interface{}) error {

	return pos.ErrorFor(tb.pos, cause, errMsg, args...)
}

func (tb tokenBuilder) String() string {
	return string(tb.out.List())
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
	if !tb.More() {
		return false, nil
	}

	have, e := tb.Peek()
	if e != nil {
		return false, e
	}

	if !f(have) {
		return false, nil
	}

	if _, e = tb.Read(); e != nil {
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

	if !tb.More() {
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
	tb.out.WriteMany(ru...)
	tb.pos.IncString(string(ru))
}

func (tb *tokenBuilder) build() token.Token {
	s := tb.String()

	rng := pos.RangeFor(tb.start, tb.pos)
	tk := token.MakeToken(tb.tt, s, rng)

	tb.start = tb.pos
	tb.tt = token.Unknown
	tb.out = inout.NewListOutput[rune]()

	return tk
}
