package token

type Token int

const (
	Unknown Token = iota
	If
	For
	Watch
	When
	E
	F
	End
)

var strMap = map[Token]string{
	If:    "if",
	For:   "for",
	Watch: "watch",
	When:  "when",
	E:     "expression",
	F:     "function",
	End:   "end",
}

func (tk Token) String() string {
	return strMap[tk]
}

func (tk Token) IsKeyword() bool {
	switch tk {
	case If, For, Watch, When, E, F, End:
		return true
	default:
		return false
	}
}
