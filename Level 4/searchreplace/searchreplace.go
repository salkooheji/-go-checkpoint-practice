package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	one := arg[1]
	two := arg[2]
	three := arg[3]

	for _, i := range one {
		if i == rune(two[0]) {
			i = rune(three[0])
		}
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
