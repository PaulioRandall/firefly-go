
# EBNF Scanning Rules

```
Char     := ? Any Unicode character ?
Letter   := ? Any Unicode character from the letter category ?
Spaace   := ? Any Unicode character from the space category except linefeed ?
Linefeed := ? Unicode linefeed character ?
```

## Comment

```
Comment := "//" (Char - Linefeed) Linefeed  
```

## Variable

```
VariableLetter := Letter | "_"
Variable       := VariableLetter {VariableLetter}  
```

## Number

```
Digit   := '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'
Integer := Digit {Digit}
Number  := Integer ["." Integer]
```

## String

```
StringEscape := '\"'
StringLetter := Letter -'"'
String       := '"' {StringLetter | StringEscape} '"'
```

## Keywords

```
Def   := "def"
If    := "if"
For   := "for"
In    := "in"
Watch := "watch"
When  := "when"
Is    := "is"
E     := "E"
F     := "F"
End   := "end"
Bool  := "true" | "false"
```

## Operators & other symbols

```
Terminator   := ";"
Assign       := "="
Comma        := ","
Colon        := ":"
Spell        := "@"

Add          := "+"
Sub          := "-"
Mul          := "*"
Div          := "/"
Mod          := "%"

LT           := "<"
GT           := ">"
LTE          := "<="
GTE          := ">="
EQU          := "=="
NEQ          := "!="

ParenOpen    := "("
ParenClose   := ")"
BraceOpen    := "{"
BraceClose   := "}"
BracketOpen  := "["
BracketClose := "]"
```
