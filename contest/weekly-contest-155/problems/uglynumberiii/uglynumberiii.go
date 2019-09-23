package nthuglynumber

import (
	"errors"
	"math"
)

//find smallest among three number
func min(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		}
		return z
	}
	if y < z {
		return y
	}
	return z
}

func nthUglyNumber(n int, a int, b int, c int) int {
	var count int
	for i := a; i < 2*int(math.Pow10(9)); i++ {
		if i < c {
			i = c
			continue
		}
		if i%a == 0 || i%b == 0 || i%c == 0 {
			count++
		}
		if count == n {
			return i
		}
	}
	panic(errors.New("No solutions"))
}
