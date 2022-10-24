
# Changes from v3 to v4

## Breaking changes

#### Function & procedure definitions

Functions and procedures will not be defined using a new keyword `def` and the
definition operator `:=` will be removed.

v3:

```
add := F(a, b) a + b

sub := P(a, b) result
	result = a - b
end
```

v4:

```
def add F(a, b) a + b

def sub P(a, b) result
	result = a - b
end
```

See `func.ff` and `proc.ff` for examples

## Non-breaking changes

#### Do block

A `do` block is introduced that allows nested scoped blocks. The primary
purpose is allow multiple statements (block) as a When case body. Currently
only a single statement is may be used.

It also enables nested blocks which have there own scope. Useful when
refactoring or for improving readability.

See `do.ff` for examples
