package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	arg := os.Args[1]

	for _, i := range arg {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {
			if i >= 'a' && i <= 'm' || i >= 'A' && i <= 'M' {
				z01.PrintRune(i + 13)
			} else {
				z01.PrintRune(i - 13)
			}
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
