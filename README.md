
# Firefly

Firefly is a simple toy scripting language aimed at replacing Bash and Python scripting in many use cases.

**Overview:**

- Fantasy themed
- Iconic syntax
- Scripting and domain specific
- Inbuilt formatter, analysis tools, and documentation server
- Easy add, remove, and modify custom functions

**These are the sorts of activities I had in mind:**

- Gluing together small programs to create a workflow
- Web API testing
- Processing files
- File watching
- Pipelining

**Some principles, rules, contraints to design and implement by:**

- Specification first
- No install
- No packaging or libraries, just scripting
- No explicit types, just scripting
- Explicit type conversions, no duck typing
- No side effects where possible
- Multiple return on functions
- Arbitrary length numbers

## Why?

What I want from a scripting tool:

- **Simplicity & minimalist:** I want scripts that are concise and to the point.
- **Readability & changability:** I want scripts that are easy to read and modify.
- **Usability:** I want to be able to execute scripts without installing a load of tools and packages first.
- **Writability**: I want scripts that are easy and quick to write but I will not sacrifice readability or changability for it.
- **Functionality**: I want to be able to perform intrinsically complex and performance sensitive functionality but I will not sacrifice simplicity for it.

[Bash](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) is core to Unix computing but the resultant code is difficult to read, understand, and maintain. Not to mention the restrictiveness in functionality.

[Python](https://en.wikipedia.org/wiki/Python_(programming_language)) gets heavy usage in scientific, mathematical, and data domains but common implementations require installation followed by a some package management tools. The design tries to capture both writability and readability but usually results in excessive idiomatic functionality being crammed into hard to read and debug one liners. Python started out with similar goals and use cases as Firefly but over time has become a bit too bloated and too systems like for my needs.

What I really want is something in between these two. I want to glue stuff together to create simple reusable scripts with more functionality than what Bash offers without the baggage required by Python. 
