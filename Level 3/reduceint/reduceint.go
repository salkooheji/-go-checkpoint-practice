package piscine

import (
	"strconv"

	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {

	x := a[0]
	for i := 1; i < len(a); i++ {
		x = f(x, a[i])
	}
	n := strconv.Itoa(x)
	for _, o := range n {
		z01.PrintRune(o)
	}
	z01.PrintRune('\n')
}
