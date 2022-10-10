
# Parse Rules

Package parser parses a series of tokens into A series of abstract syntax trees

- Words with capital first letter with remaining being lowercase are tokens

```
STATEMENT  := (ASSIGNMENT | BRANCH) Terminator

BRANCH     := If CONDITION Terminator {STATEMENT} Terminator End
ASSIGNMENT := Variable {Comma Variable} Assign EXPRESSION {Comma EXPRESSION}

CONDITION  := EXPRESSION

EXPRESSION := LITERAL

LITERAL    := True | False | Number | String
```
