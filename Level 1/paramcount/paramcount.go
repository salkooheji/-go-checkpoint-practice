package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	a := len(os.Args[1:])
	z01.PrintRune(rune(a + 48))
	z01.PrintRune('\n')
}
