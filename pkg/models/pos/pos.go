package pos

import (
	"fmt"
)

type Pos struct {
	Offset int
	Line   int // index
	Col    int // index
}

func At(offset, line, col int) Pos {
	return Pos{
		Offset: offset,
		Line:   line,
		Col:    col,
	}
}

func IncRune(p Pos, ru rune) Pos {
	p.IncRune(ru)
	return p
}

func IncString(p Pos, s string) Pos {
	p.IncString(s)
	return p
}

func (p *Pos) IncRune(ru rune) {
	p.Offset++

	if ru == '\n' {
		p.Line++
		p.Col = 0
	} else {
		p.Col++
	}
}

func (p *Pos) IncString(s string) {
	for _, ru := range s {
		p.IncRune(ru)
	}
}

// ***** RETIRE *****

func (p Pos) String() string {
	return fmt.Sprintf("Offset=%d Line=%d Col=%d", p.Offset, p.Line, p.Col)
}

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
		At(offset, line, col),
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
