package sqrtx

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	lo, hi := 0, x

	var mi int
	for {
		mi = (hi + lo) / 2
		if mi == lo {
			return mi
		}
		r := mi * mi
		if int(r) == x {
			return mi
		} else if int(r) < x {
			lo = mi
		} else {
			hi = mi
		}
	}
}
