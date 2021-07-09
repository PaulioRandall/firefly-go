# Firefly

Firefly will be my third attempt at a programming language. Alongside the experience I gained in the last two attempts I aim to put heavier focus on specification and interfaces.

Firefly is a toy programming language built on the following ideas:

- Fantasy themed
- Specification first
- No install
- Modification over extension

## Versions & Ramblings

**v0.1.0 (Current)**

This version allows you to perform arithmetic operations on a single line which print to standard output.

**v0.2.0**

For this version I aim to add an arbitrary precise decimal number type or something close to it. At the moment divinding an odd number by an even will round down, e.g. `5 / 2 = 2`. I will also be updating integers to be arbitrarly precise.

Type conversions will be explicit. This was a long and hard decision, but I decided that forcing explicit type conversions avoided subtly wrong output errors. This decision was greatly influenced by reading the discussions on type conversions in Go.

Adding a `.0` to the end of any literal integer is the same as using the `float` type conversion and is the recommended way of specifying floats. Type conversions are more for named values (on the fly defined constants) such as `x`.

## Definitions

| Term | Description |
| :--- | :--- |
| _Scroll_ | The text or file containing the instructions for a program or script |

## Examples

This scroll adds together two numbers and prints the result to standard output, usually the console.

```
7 + 1
```

This scroll is a little more advanced. It performs two mathematical operations and prints them on their own line with an empty line in between. The first expression involves float types while the second involves integers.

```
2.0 * (10.3 - 6.3)

9 / 4
```

This scroll shows off named values and assignment statements. Assignment statements have the name on the left, an assignment operator `:=` in the middle, and the expression which defines the value referred to the name on the right. The first two statements create named values which are then used in the third expression. Assignment statements only print a linefeed when executed unlike expressions. The fifth line does the same thing as the first three but without named values.

```
x := 5 - 1
y := 2 * 2
x + y

(5 - 1) + (2 * 2)
```

This scroll demonstrates type conversion. The output of the second statement will be `2.5`. If `x` was not converted to a float then a compile error would occur as there is no implicit type conversion, not even between integers and floats.

```
x := 5
float(x) / 2.0
```

This statment converts the number within to an int. In this case it will round down to `2`. Note that attempting to type convert a number to it's current type will produce the same result as not having the type conversion at all.

```
int(2.8)
```

This statement will output `2.0` because the operation within the type conversion has two interger operands. The type conversion happens after the inner expresssion has been executed.

```
float(5 / 2)
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
