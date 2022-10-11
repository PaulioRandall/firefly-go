package scanner

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

var ErrNotFound = err.New("Unknown symbol")

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
		return rune(0), err.Wrap(e, "Failed to read rune")
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
		return false, err.Wrap(e, "Failed to accept rune")
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
		return err.Wrap(inout.EOF, "Unexpected EOF")
	}

	found, e := tb.acceptFunc(f)
	if e != nil {
		return err.Wrapf(e, errMsg, args...)
	}

	if !found {
		return err.Wrapf(ErrNotFound, errMsg, args...)
	}

	return nil
}

func (tb *tokenBuilder) build() token.Token {
	s := tb.String()

	tk := token.MakeToken(tb.tt, s, tb.start, tb.r.Where())

	tb.start = tb.r.Where()
	tb.tt = token.Unknown
	tb.w = inout.NewListWriter[rune]()

	return tk
}
