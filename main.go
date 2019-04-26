package main

import (
	"os"

	"github.com/shozawa/monkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
