
# Firefly

Firefly is a toy programming language with the following dream characteristics:

- Fantasy themed
- Specification first
- Domain specific
- No install
- Modification over extension
- No packaging or libraries, just scripting
- No explicit types, just scripting
- Explicit type conversions, no duck typing
- Arbitrary length numbers
- Multiple return on functions
- No side effects where possible
- Iconic syntax
- Inbuilt formatter
- Inbuilt documentation server
- Easy add, remove, and modify custom functions

Targeting the following domains and not considering other domains:

- Simple scripting, something one level up from bash
- Gluing together small programs to create a workflow
- Web API testing
- Processing files
- File watcher

This specification defines `v3: Gloryhammer`.

## Syntax

### Loops

There are two schemas for looping:

```
for <index>, <item> in <list> {
	<statements>
}
```

```
for <initial statement>; <condition>; <between iteration statement> {
	<statements>
}
```

- Loops 10 times
- On each loop:
  - Assign i the current index and n as i+1
  - Print "i : n" where i and n are replaced by their current assigned values

```
for i, n in @numbers(0, 10) {
	@print(i, " : ", n)
}
```

- Loop until the condition is false

```
condition = true
for condition {
	condition <- false
}


// Can be rewritten as
for condition = true; condition {
	condition <- false
}
```

### Functions

- A function
- Can have side effects
- Return arguments are specified upfront
- There is no return keyword 

```
divide <- F(a, b) result, error {
	when b == 0 {
		true:  error  <- "Can't divide by zero"
		false: result <- a / b
	}
}

result, error <- divide(4, 2)
```

### Expression functions

- Basically an expression that can be inlined
- Always pure, i.e. can't have side effects
- Always a one liner
- Always returns a single value

```
add <- E(a, b) a + b

result <- add(4, 2)
```

### When

- Basically a match block
- May or may not have a subject
- Executes the first matching case then exits the when
- Cases can be:
  - a static value
  - a variable
  - a condition (expression)
- Each case statement can be inline or a block

```
when v {
	0: @println("Is 0")
	1: @println("Is 1")
	2: @println("Is 2")
}

when {
	v == 0: @println("Is 0")
	v == 1: @println("Is 1")
	v == 2: @println("Is 2")
}

when v {
	v == 0: @println("Is 0")
	1:      @println("Is 1")
	2:      @println("Is 2")
	v > 2:  @println("Greater than 2")
	(true):   {
		@println("Negative")
	}
}
```
