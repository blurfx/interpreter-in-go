package main

import (
	"interpreter-in-go/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
