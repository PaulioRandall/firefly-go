package token

type Token int

const (
	Unknown Token = iota
	If
)

var strMap = map[Token]string{
	If: "if",
}

func (tk Token) String() string {
	return strMap[tk]
}

func (tk Token) IsKeyword() bool {
	switch tk {
	case If:
		return true
	default:
		return false
	}
}
