package token

type TokenType int
type TypeMetadata struct {
	Type   TokenType
	Name   string
	Symbol string
}

const (
	Unknown TokenType = iota

	Newline    // '\n'
	Terminator // ;

	Space
	Comment

	Ident
	Bool
	Number
	String

	Def
	If
	For
	In
	Watch
	When
	Is
	Func
	Proc
	End

	Assign // =
	Comma  // ,
	Colon  // :
	Spell  // @

	Add // +
	Sub // -
	Mul // *
	Div // /
	Mod // %

	Lt  // <
	Mt  // >
	Lte // <=
	Mte // >=
	Equ // ==
	Neq // !=

	And // &&
	Or  // ||

	ParenOpen  // (
	ParenClose // )

	BraceOpen  // {
	BraceClose // }

	BracketOpen  // [
	BracketClose // ]
)

func typeMeta(typ TokenType, name, symbol string) TypeMetadata {
	return TypeMetadata{
		Type:   typ,
		Name:   name,
		Symbol: symbol,
	}
}

var Metadata = map[TokenType]TypeMetadata{
	Unknown: typeMeta(Unknown, "Unknown (zero value)", ""),

	// Terminators
	Newline:    typeMeta(Newline, "Newline", "\n"),
	Terminator: typeMeta(Terminator, "Terminator", ";"),

	// Redundant
	Space:   typeMeta(Space, "Space", ""),
	Comment: typeMeta(Comment, "Comment", ""),

	// Terms
	Ident:  typeMeta(Ident, "Identifier", ""),
	Bool:   typeMeta(Bool, "Bool", ""),
	Number: typeMeta(Number, "Number", ""),
	String: typeMeta(String, "String", ""),

	// Keywords
	Def:   typeMeta(Def, "Define", "def"),
	If:    typeMeta(If, "If", "if"),
	For:   typeMeta(For, "For", "for"),
	In:    typeMeta(In, "In", "in"),
	Watch: typeMeta(Watch, "Watch", "watch"),
	When:  typeMeta(When, "When", "when"),
	Is:    typeMeta(Is, "Is", "is"),
	Func:  typeMeta(Func, "Function", "F"),
	Proc:  typeMeta(Proc, "Procedure", "P"),
	End:   typeMeta(End, "End", "end"),

	// Operators
	Assign: typeMeta(Assign, "Assignment", "="),
	Comma:  typeMeta(Comma, "Comma", ","),
	Colon:  typeMeta(Colon, "Colon", ":"),
	Spell:  typeMeta(Spell, "Spell", "@"),

	// Arithmetic operators
	Add: typeMeta(Add, "Add", "+"),
	Sub: typeMeta(Sub, "Subtract", "-"),
	Mul: typeMeta(Mul, "Multiply", "*"),
	Div: typeMeta(Div, "Divide", "/"),
	Mod: typeMeta(Mod, "Remainder", "%"),

	// Comparison operators
	Lt:  typeMeta(Lt, "Less than", "<"),
	Mt:  typeMeta(Mt, "More than", ">"),
	Lte: typeMeta(Lte, "Less than or equal", "<="),
	Mte: typeMeta(Mte, "More than or equal", ">="),
	Equ: typeMeta(Equ, "Equal", "=="),
	Neq: typeMeta(Neq, "Not equal", "!="),

	// Boolean operators
	And: typeMeta(And, "And", "&&"),
	Or:  typeMeta(Or, "Or", "||"),

	// Parentheses
	ParenOpen:    typeMeta(ParenOpen, "Paren open", "("),
	ParenClose:   typeMeta(ParenClose, "Paren close", ")"),
	BraceOpen:    typeMeta(BraceOpen, "Brace open", "{"),
	BraceClose:   typeMeta(BraceClose, "Brace close", "}"),
	BracketOpen:  typeMeta(BracketOpen, "Bracket open", "["),
	BracketClose: typeMeta(BracketClose, "Bracket close", "]"),
}

func (tt TokenType) String() string {
	return Metadata[tt].Name
}

func (tt TokenType) Name() string {
	return Metadata[tt].Name
}

func (tt TokenType) Symbol() string {
	return Metadata[tt].Symbol
}

func (tt TokenType) Precedence() int {
	switch tt {
	case Mul, Div, Mod:
		return 6
	case Add, Sub:
		return 5
	case Lt, Mt, Lte, Mte:
		return 4
	case Equ, Neq:
		return 3
	case And:
		return 2
	case Or:
		return 1
	default:
		return 0
	}
}
