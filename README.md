# Firefly

Firefly will be my third attempt at a programming language. Alongside the experience I gained in the last two attempts I aim to put heavier focus on specification and interfaces.

Firefly is a toy programming language built on the following ideas:

- Fantasy themed
- Specification first
- No install
- Modification over extension

## Versions & Ramblings

As of v0.1.0 you can perform arithmetic operations on a single line which print to standard output. For v0.2.0 I aim to add an arbitrary precise decimal number type or something close to it. At the moment divinding an odd number by an even will round down, e.g. `5 / 2 = 2`.

I have to decide if this new type will replace the integer type or not. The decision is a battle between my cravings for a minimalist design and my yearnings for integrity and precise engineering. If I go for separate types then the integer type will be upgraded from 64 bits to an arbitary precise integer or something close to it.

This also introduces another dilemma, type coercion. I feel the simplist solution is to implicitly upgrade an integer to a decimal when performing a mixed calculation such that the use of a decimal anywhere in an expression will always produce a decimal result. Since systems programming is not an intended use of this language --I'm fully aware I have yet to specify where Firefly sits on the high level language spectrum-- I don't see any reason to force explicit type conversion in such cases. However, downgrading from decimals to integers will need to be explicit.

I thought downgrading would not be needed just yet as we are only printing the results of each statement to standard output. After a moments thought I realised the decision depends on whether the same number in each type outputs in the same manner or not.

If so, I have the choice of syntax for converting types. The two immediate options that come to mind are C/C++/Java style `(int) n` and Go style `int(n)`. I will need to research alternative approaches and maybe have a go at devising a new one. But of the two, I prefer the Go style since it's clearer what value is being converted in some cases. E.g. using C style I've seen code like `(int) 1.1 * 2.2` which drips with ambiguity, at least from a readability perspective. I want to steer strongly towards iconic code, i.e. code that is unambiguous and familiar to those who have used other imperative languages.

Hiding here is another subtle issue. What happens when you attempt to convert an integer to an integer? Or attempt to convert any value to its current type? I see no reason to raise an error even if the instruction is redundant. Furthermore, it could be used for security, ensuring the value is of the type wanted, and maybe for showing intention through code.

Those last two points suggest an explicit upgrade might be useful too. If Go style is used then this would probably look like `float(n)`. I've used the term `decimal` so far but I think `float` would be more iconic and accurate. Particularly if following the general definitions of the two terms in computer science, i.e. `decimal` infers greater accuracy but lower performance than `float`.

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
