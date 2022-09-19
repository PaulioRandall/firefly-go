package token

type TokenType int

const (
	Unknown TokenType = iota

	Var

	_keywords_begin
	If
	For
	In
	Watch
	When
	E
	F
	End
	True
	False
	_keywords_end

	_operators_begin

	Ass        // =
	Def        // :=
	Terminator // ;
	Comma      // ,

	_arith_begin
	Add // +
	Sub // -
	Mul // *
	Div // /
	Mod // %
	_arith_end

	_cmp_begin
	LT  // <
	GT  // >
	LTE // <=
	GTE // >=
	EQU // ==
	NEQ // !=
	_cmp_end

	_paren_begin
	ParenOpen    // (
	ParenClose   // )
	BraceOpen    // {
	BraceClose   // }
	BracketOpen  // [
	BracketClose // ]
	_paren_end
	_operators_end
)

var symbolMap = map[TokenType]string{
	Var:          "",
	If:           "if",
	For:          "for",
	In:           "in",
	Watch:        "watch",
	When:         "when",
	E:            "E",
	F:            "F",
	End:          "end",
	True:         "true",
	False:        "false",
	Ass:          "=",
	Def:          ":=",
	Terminator:   ";",
	Comma:        ",",
	Add:          "+",
	Sub:          "-",
	Mul:          "*",
	Div:          "/",
	Mod:          "%",
	LT:           "<",
	GT:           ">",
	LTE:          "<=",
	GTE:          ">=",
	EQU:          "==",
	NEQ:          "!=",
	ParenOpen:    "(",
	ParenClose:   ")",
	BraceOpen:    "{",
	BraceClose:   "}",
	BracketOpen:  "[",
	BracketClose: "]",
}

func (tt TokenType) IsKeyword() bool {
	return tt > _keywords_begin && tt < _keywords_end
}

func (tt TokenType) IsOperator() bool {
	return tt > _operators_begin && tt < _operators_end
}

func (tt TokenType) String() string {
	return symbolMap[tt]
}

func IdentifyWordType(s string) TokenType {
	for tt, symbol := range symbolMap {
		if tt.IsKeyword() && s == symbol {
			return tt
		}
	}

	return Var
}

func IdentifyOperatorType(s string) TokenType {
	for tt, symbol := range symbolMap {
		if tt.IsOperator() && s == symbol {
			return tt
		}
	}

	return Unknown
}
