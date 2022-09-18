package token

type Pos struct {
	Idx  int
	Line int
	Col  int // Index on line
}

type Range struct {
	Start Pos
	End   Pos // Exclusive
}

func MakePos(idx, line, col int) Pos {
	return Pos{
		Idx:  idx,
		Line: line,
		Col:  col,
	}
}

func MakeRange(start, end Pos) Range {
	return Range{
		Start: start,
		End:   end,
	}
}

func (p *Pos) Inc(ru rune) {
	p.Idx++

	if ru == '\n' {
		p.Line++
		p.Col = 0
	} else {
		p.Col++
	}
}
