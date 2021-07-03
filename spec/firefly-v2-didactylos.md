
# Firefly

Firefly is a toy programming language built on the following ideas:

- Fantasy themed
- Specification first
- No install
- Modification over extension

This specification defines `v2: Didactylos`.

## Definitions

| Term | Description |
| :--- | :--- |
| _Scroll_ | The text or file containing the instructions for a program or script |

## Examples

This scroll adds together two numbers and prints the result to standard output, usually the console.

```
7 + 1
```

This scroll is a little more advanced. It performs two mathematical operations and prints them on their own line with an empty line in between.

```
2 * (10.3 - 6.3)

9 / 4
```

This scroll shows off named values and assignment statements. Assignment statements have the name on the left, an assignment operator `:=` in the middle, and the expression which defines the value referred to the name on the right. The first two statements create named values which are then used in the third expression. Assignment statements only print a linefeed when executed unlike expressions. The fifth line does the same thing as the first three but without named values.

```
x := 5 - 1
y := 2 * 2
x + y

(5 - 1) + (2 * 2)
```

This scroll demonstrates type conversion. The output of the first statement will be `2.5` since any expression containing a float will force the other operand in the operation to always upgrade from an int to a float, if not already. The second statement generates the same output but more concise and probably more intuitive in this example. 

```
float(5) / 2
5.0 / 2
```

This statment converts the number within to an int. In this case it will round down to `2`. Note that attempting to type convert a number to it's current type will produce the same result as not having the type conversion at all.

```
int(2.8)
```

This statement will output `2.0` because the operation within the type conversion has two interger operands. The type conversion happens after the inner expresssion has been completed.

```
float(5 / 2)
```

This final scroll will output `7.5` because the right hand sub expression will produce `3` and not `3.3333...`. Both operand types are integers, hence the result will be an integer. As a sub expression, it gets processed in isolation from the other left hand sub expression before being used in the multiplication.

```
(5.0 / 2) * (10 / 3)
```

## EBNF

A set of production rules that define the syntax of a Firefly program.

```
PROGRAM := { STATEMENT }

STATEMENT  := EXPRESSION | ASSIGNMENT
ASSIGNMENT := NAMED_VALUE ":=" EXPRESSION

EXPRESSION := EXPRESSION ARITHMETIC_OPERATOR EXPRESSION
EXPRESSION := SUB_EXPRESSION
EXPRESSION := OPERAND

SUB_EXPRESSION := "(" EXPRESSION ")"

ARITHMETIC_OPERATOR := "+" | "-" | "*" | "/"
OPERAND             := INTEGER | NAMED_VALUE | TYPE_CONVERSION

TYPE_CONVERSION := TYPE_NAME SUB_EXPRESSION
TYPE_NAME       := "int" | "float"

NAMED_VALUE := LETTER { LETTER }
LETTER      := *Any lowercase character from the English alphabet*

FLOAT   := INTEGER "." DIGIT { DIGIT }
INTEGER := NON_ZERO_DIGIT { DIGIT }
INTEGER := ZERO_DIGIT

DIGIT          := ZERO_DIGIT | NON_ZERO_DIGIT
ZERO_DIGIT     := "0"
NON_ZERO_DIGIT := "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```

## Types

| Name | Description | Examples |
| :--- | :--- | :--- |
| _Integer_ | An arbitary precise integer | `0`, `21`, `123456` |
| _Float_ | An arbitary precise floating point number | `123456.0`, `123.456`, `0.123456` |

## Operators

| Symbol | Precedence | Description | 
| :---: | :---: | :--- |
| `(` `)` | 3 | Parentheses for encapsulating sub expressions |
| `*` | 2 | Numeric [multiplication](https://en.wikipedia.org/wiki/Multiplication) |
| `/` | 2 | Numeric [division](https://en.wikipedia.org/wiki/Division_(mathematics)) |
| `+` | 1 | Numeric [addition](https://en.wikipedia.org/wiki/Addition) |
| `-` | 1 | Numeric [subtraction](https://en.wikipedia.org/wiki/Subtraction) |

## Rules

### Spaces & newlines

- Whitespace is ignored and used only to separate non-terminal symbols
- Linefeeds `\n` are considered statement terminators and not whitespace
- Lines with only whitespace or no characters are called empty statements

### Statements

- Each statement
	- starts at the beginning of its line
	- is terminated by a linefeed `\n`
	- can either be assignment or expression
	- is executed in the order specified within a scroll
- Upon execution of a statement
	- [empty statement] a linefeed is printed to standard output
	- [expression] the result is printed before a linefeed to standard output
	- [assignment] the result of its expression is bound or mapped to its name and stored for use in further statements

### Errors

- On error, it is printed to the error output and the program will immediately terminate
