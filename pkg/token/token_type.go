package token

type TokenType int

const (
	Unknown TokenType = iota

	Var

	_keywords_begin
	If
	For
	Watch
	When
	E
	F
	End
	True
	False
	_keywords_end

	_operators_begin
	Add // +
	Sub // -
	Mul // *
	Div // /
	Mod // %
	LT  // <
	GT  // >
	LTE // <=
	GTE // >=
	EQU // ==
	NEQ // !=
	ASS // =
	DEF // :=
	_operators_end
)

var syntaxMap = map[TokenType]string{
	If:    "if",
	For:   "for",
	Watch: "watch",
	When:  "when",
	E:     "E",
	F:     "F",
	End:   "end",
	True:  "true",
	False: "false",
	Var:   "",
	Add:   "+",
	Sub:   "-",
	Mul:   "*",
	Div:   "/",
	Mod:   "%",
	LT:    "<",
	GT:    ">",
	LTE:   "<=",
	GTE:   ">=",
	EQU:   "==",
	NEQ:   "!=",
	ASS:   "=",
	DEF:   ":=",
}

func (tt TokenType) IsKeyword() bool {
	return tt > _keywords_begin && tt < _keywords_end
}

func (tt TokenType) IsOperator() bool {
	return tt > _operators_begin && tt < _operators_end
}

func (tt TokenType) String() string {
	return syntaxMap[tt]
}

func IdentifyWordType(s string) TokenType {
	for tt, symbol := range syntaxMap {
		if tt.IsKeyword() && s == symbol {
			return tt
		}
	}

	return Var
}

func IdentifyOperatorType(s string) TokenType {
	for tt, symbol := range syntaxMap {
		if tt.IsOperator() && s == symbol {
			return tt
		}
	}

	return Unknown
}
