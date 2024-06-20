package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	count1 := 0
	count2 := 0

	for count1 < len(str1) && count2 < len(str2) {
		if str1[count1] == str2[count2] {
			count1++
		}
		count2++
	}
	if count1 == len(str1) {
		for _, i := range str1 {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
