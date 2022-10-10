package scanner

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/utilities/inout"
	"github.com/PaulioRandall/firefly-go/utilities/pos"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var ErrNotFound = errors.New("Unknown symbol")

type tokenBuilder struct {
	r     inout.RuneReader
	w     inout.ListWriter[rune]
	start pos.Pos
	tt    token.TokenType
}

func newTokenBuilder(r ReaderOfRunes) tokenBuilder {
	return tokenBuilder{
		r: inout.NewRuneReader(r),
		w: inout.NewListWriter[rune](),
	}
}

func (tb *tokenBuilder) err(cause error, errMsg string, args ...interface{}) error {
	return pos.ErrorFor(tb.r.Where(), cause, errMsg, args...)
}

func (tb tokenBuilder) String() string {
	return string(tb.w.List())
}

func (tb tokenBuilder) More() bool {
	return tb.r.More()
}

func (tb *tokenBuilder) Peek() (rune, error) {
	return tb.r.Peek()
}

func (tb *tokenBuilder) Read() (rune, error) {
	ru, e := tb.r.Read()
	if e != nil {
		return rune(0), tb.err(e, "[tokenBuilder] Failed to read rune")
	}

	tb.w.Write(ru)
	return ru, nil
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

	tb.w.Write(have)
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

	if !tb.r.More() {
		return inout.EOF
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

func (tb *tokenBuilder) build() token.Token {
	s := tb.String()

	rng := pos.RangeFor(tb.start, tb.r.Where())
	tk := token.MakeToken(tb.tt, s, rng)

	tb.start = tb.r.Where()
	tb.tt = token.Unknown
	tb.w = inout.NewListWriter[rune]()

	return tk
}
