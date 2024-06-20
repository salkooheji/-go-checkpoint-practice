package main

import (
   "os" 
    "github.com/01-edu/z01"
    )
    
func main () {
    
    if len(os.Args) != 2 {
        return
    }
    
    args := os.Args[1]
    LastWord := ""
    
    if args == " " {
        return
    }
    
    for i := len(args) - 1 ; i >= 0 ; i-- {
        if args[i] != ' ' {
            LastWord = string(args[i]) + LastWord
        } else if LastWord != "" {
            break
        }
    }
   for _, j := range LastWord {
       z01.PrintRune(j)
   }
   z01.PrintRune('\n') 
}