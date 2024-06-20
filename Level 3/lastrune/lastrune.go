package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
var i rune
for _, a := range s {
    i = a
}
return i
}
