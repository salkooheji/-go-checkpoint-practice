package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) == 1 {
		return
	}

	args := os.Args[1]

	for _, i := range args {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
