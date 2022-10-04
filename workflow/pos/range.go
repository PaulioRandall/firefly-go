package pos

import (
	"fmt"
)

type Range struct {
	From Pos
	To   Pos // exclusive
}

func MakeRange(from, to Pos) Range {
	return Range{
		From: from,
		To:   to,
	}
}

func (r *Range) IncString(s string) {
	r.From = r.To
	r.To.IncString(s)
}

func (r Range) String() string {
	return fmt.Sprintf("from { %v } to { %v }", r.From, r.To)
}
