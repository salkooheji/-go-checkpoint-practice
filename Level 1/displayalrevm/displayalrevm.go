package main

import (
    "github.com/01-edu/z01"
    )
    
func main () {
    a := "zYxWvUtSrQpOnMlKjIhGfEdCbA"
    
    for _, i := range a {
        z01.PrintRune(i)
    }
z01.PrintRune('\n')
}