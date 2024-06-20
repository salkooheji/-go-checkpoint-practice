package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) == 1 {
		return
	}

	args := os.Args[1:]

	Final := args[len(args)-1]

	for _, i := range Final {
		z01.PrintRune(i)
	}

	z01.PrintRune('\n')

}
