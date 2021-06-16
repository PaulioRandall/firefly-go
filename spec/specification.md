
# Firefly

Firefly is a toy programming language built on the following ideas:

- Fantasy themed
- Specification first
- No install
- Easily modifiable and extendable (to the extent that is reasonable possible)

This specification defines V1: 鳥取県:鳥取市 (Tottori-ken:Tottori-shi).

## Definitions

| Term | Description |
| :--- | :--- |
| _Scroll_ | The text or file containing the instructions for a program or script |

## Examples

This scroll adds together two numbers and prints the result to standard output, usually the console.

```
7 + 1
```

This scroll is a little more advanced. It performs two mathematical operations and prints them on their own line.

```
2 * (10 - 6)
24 / 4
```

## EBNF

A set of production rules that define the syntax of a Firefly program.

```
PROGRAM := { STATEMENT }

STATEMENT := { EXPRESSION }

EXPRESSION := EXPRESSION OPERATOR EXPRESSION
EXPRESSION := OPERAND

OPERATOR := "+" | "-" | "*" | "/"
OPERAND  := DIGIT

DIGIT := "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```

## Types

| Name | Description | Examples |
| :--- | :--- | :--- |
| _Number_ | A non-negative integer | `0`, `21`, `123456` |

## Operators

| Symbol | Precedence | Description | 
| :---: | :---: | :--- |
| `(` `)` | 3 | Parentheses for encapsulating expressions |
| `*` | 2 | Numeric [multiplication](https://en.wikipedia.org/wiki/Multiplication) |
| `/` | 2 | Numeric [division](https://en.wikipedia.org/wiki/Division_(mathematics)) |
| `+` | 1 | Numeric [addition](https://en.wikipedia.org/wiki/Addition) |
| `-` | 1 | Numeric [subtraction](https://en.wikipedia.org/wiki/Subtraction) |

## Rules

1. Whitespace is ignored and used only to separate non-terminal symbols
2. Linefeeds `\n` are not considered whitespace
3. Each instruction starts at the first non-whitespace symbol on a line
4. Each instruction is terminated by a linefeed `\n`, the linefeed is consumed in the process 
5. Each instruction is a mathematical calculation
6. The result of each instruction is printed, on its own line, to the standard output
7. Instructions are performed in the order specified within a scroll
8. If an error is encountered it is printed to the error output and the program will immediately terminate
