// Program firefly is a toy programming language built on the following ideas:
//
//		Fantasy themed
//		Specification first
//		No install
//		Easily modifiable and extendable (to the extent that is reasonable possible)
//
// This is an implementation of version 1, `Interesting times`. The
// language specification is available some where in source.
package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/firefly-go/firefly/firefly"
)

func main() {
	file := getFilenameArg()
	run(file)
}

func getFilenameArg() string {
	args := os.Args
	if len(args) < 2 {
		panic("expected filename argument")
	}
	return args[1]
}

func run(file string) {
	e := firefly.RunFile(file)
	if e != nil {
		panic(e)
	}
}

func printTaskList() {
	fmt.Print(`
Hello, firefly!


	NEXT: Check spec is Left, Leftmost, and lookahead only by 1, aka LL(1)
			+ Impose a rule that the grammer stay LL(1) for future changes
	https://stackoverflow.com/questions/8496642/how-to-identify-whether-a-grammar-is-ll1-lr0-or-slr1
`)
}
