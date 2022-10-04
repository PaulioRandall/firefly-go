package pos

import (
	"fmt"
)

type Range struct {
	From Pos
	To   Pos // exclusive
}

func RangeFor(from, to Pos) Range {
	return Range{
		From: from,
		To:   to,
	}
}

func RangeForString(from Pos, s string) Range {
	rng := Range{
		From: from,
		To:   from,
	}

	rng.ShiftString(s)
	return rng
}

func RawRangeForString(offset, line, col int, s string) Range {
	return RangeForString(
		PosAt(offset, line, col),
		s,
	)
}

func (r *Range) ShiftString(s string) {
	r.From = r.To
	r.To.IncString(s)
}

func (r Range) String() string {
	return fmt.Sprintf("from { %v } to { %v }", r.From, r.To)
}
