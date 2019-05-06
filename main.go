package main

import (
	"fmt"
	"os"

	"github.com/shozawa/monkey/interpreter"
	"github.com/shozawa/monkey/repl"
)

func main() {
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("can't open file: %q\n", os.Args[1])
		}
		defer file.Close()
		interpreter.Execute(file)
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
