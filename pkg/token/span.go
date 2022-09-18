package token

type Place int
type Span struct {
	Start     Place
	End       Place // Exclusive
	Line      Place
	Len       Place
	LineStart Place
	LineEnd   Place // Exclusive
}

func MakeSpan(start, end, line, lineStart Place) Span {
	length := end - start

	return Span{
		Start:     start,
		End:       end,
		Len:       length,
		Line:      line,
		LineStart: lineStart,
		LineEnd:   lineStart + length,
	}
}
