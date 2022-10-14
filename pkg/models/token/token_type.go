package token

type TokenType int

const (
	Unknown TokenType = iota

	Newline    // '\n'
	Terminator // ;
	Identifier

	_redundant_begin
	Space
	Comment
	_redundant_end

	_keywords_begin
	If
	For
	In
	Watch
	When
	Is
	Expr // TODO: Rename as Func
	Proc
	End
	_literal_begin
	True
	False
	_keywords_end
	Number
	String
	_literal_end

	_operators_begin
	Assign // =
	Define // :=
	Comma  // ,
	Colon  // :
	Spell  // @

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

var nameMap = map[TokenType]string{
	Newline:      "Newline",
	Terminator:   "Terminator",
	Identifier:   "Identifier",
	Space:        "Space",
	Comment:      "Comment",
	Number:       "Number",
	String:       "String",
	If:           "If",
	For:          "For",
	In:           "In",
	Watch:        "Watch",
	When:         "When",
	Is:           "Is",
	Expr:         "Expression",
	Proc:         "Procedure",
	End:          "End",
	True:         "True",
	False:        "False",
	Assign:       "Assign",
	Define:       "Define",
	Comma:        "Comma",
	Colon:        "Colon",
	Spell:        "Spell",
	Add:          "Add",
	Sub:          "Sub",
	Mul:          "Mul",
	Div:          "Div",
	Mod:          "Mod",
	LT:           "Less Than",
	GT:           "More Than",
	LTE:          "Less Than Equal",
	GTE:          "More Than Equal",
	EQU:          "Equal",
	NEQ:          "Not Equal",
	ParenOpen:    "Paren Open",
	ParenClose:   "Paren Close",
	BraceOpen:    "Brace Open",
	BraceClose:   "Brace Close",
	BracketOpen:  "Bracket Open",
	BracketClose: "Bracket Close",
}

func (tt TokenType) String() string {
	return nameMap[tt]
}

func IsLiteral(tt TokenType) bool {
	return tt > _literal_begin && tt < _literal_end
}

func IsRedundant(tt TokenType) bool {
	return tt > _redundant_begin && tt < _redundant_end
}

func IsKeyword(tt TokenType) bool {
	return tt > _keywords_begin && tt < _keywords_end
}

func IsOperator(tt TokenType) bool {
	return tt > _operators_begin && tt < _operators_end
}

func Operators() map[TokenType]string {
	return filter(func(tt TokenType, _ string) bool {
		return IsOperator(tt)
	})
}

func filter(f func(TokenType, string) bool) map[TokenType]string {
	res := map[TokenType]string{}

	for tt, symbol := range nameMap {
		if f(tt, symbol) {
			res[tt] = symbol
		}
	}

	return res
}
