package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/firefly-go/pkg/workflow"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/debug"
)

func main() {
	runWorkflow()
}

func runWorkflow() {
	file := getFile()
	if file == "" {
		fmt.Println("Expected argument: path to scroll")
		return
	}

	exitCode, e := workflow.RunFile(file)

	if e != nil {
		debug.Println(e)
		return
	}

	fmt.Println("Exit: ", exitCode)
}

func getFile() string {
	if len(os.Args) < 2 {
		return ""
	}
	return os.Args[1]
}
