package sqrtx

import (
	"math"
)

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	lo, hi := 0, x+1

	var mi int
	for lo < hi {
		mi = (hi + lo) / 2
		r := math.Pow(float64(mi), 2)
		if int(r) == x {
			return mi
		} else if int(r) < x {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return mi
}
