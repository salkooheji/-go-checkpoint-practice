package main

import (
    "github.com/01-edu/z01"
    )
    
func main () {
    a := "aBcDeFgHiJkLmNoPqRsTuVwXyZ"
    
    for _, i := range a {
        z01.PrintRune(i)
    }
z01.PrintRune('\n')
}