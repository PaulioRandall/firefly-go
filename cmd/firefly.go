package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/firefly-go/pkg/firefly"
)

func main() {
	file := getFilenameArg()
	run(file)
}

func getFilenameArg() string {
	args := os.Args
	if len(args) < 2 {
		panic("Expected filename argument")
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

NEXT: Replace Putback with Peek in readers to be more inline with LL(1) parsers
NEXT: Check spec is Left, Leftmost, and lookahead only by 1, aka LL(1)
			+ Impose rule that the grammer stay LL(1) 
https://stackoverflow.com/questions/8496642/how-to-identify-whether-a-grammar-is-ll1-lr0-or-slr1

NEXT: Add godoc support to godo
NEXT: Write documentation for pkgs
	- token
	- scanner
	- grouper
	- cleaner
	- ast
	- parser
	- runner
`)
}
