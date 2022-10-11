package pos

import (
	"fmt"
)

type Wherer interface {
	Where() (from, to Pos)
}

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

func EndAt(offset, line, col int, s string) Pos {
	p := At(offset, line, col)
	p.IncrementBy(s)
	return p
}

func EndOf(from Pos, s string) Pos {
	from.IncrementBy(s)
	return from
}

func RangeFor(from Pos, s string) (Pos, Pos) {
	return from, EndOf(from, s)
}

func RawRangeFor(offset, line, col int, s string) (Pos, Pos) {
	from := At(offset, line, col)
	return from, EndOf(from, s)
}

func (p *Pos) Increment(ru rune) {
	p.Offset++

	if ru == '\n' {
		p.Line++
		p.Col = 0
	} else {
		p.Col++
	}
}

func (p *Pos) IncrementBy(s string) {
	for _, ru := range s {
		p.Increment(ru)
	}
}

func (p Pos) String() string {
	return fmt.Sprintf("Offset=%d Line=%d Col=%d", p.Offset, p.Line, p.Col)
}

func RangeString(from, to Pos) string {
	return fmt.Sprintf("from { %v } to { %v }", from, to)
}

func WhereString(wherer Wherer) string {
	from, to := wherer.Where()
	return fmt.Sprintf("from { %v } to { %v }", from, to)
}
