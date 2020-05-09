package validperfectsquare

func isPerfectSquare(num int) bool {

	lo, hi := 0, num

	for lo < hi {
		mi := (lo + hi) / 2
		pow := mi * mi
		if pow == num {
			return true
		} else if pow > num {
			hi = mi
		} else {
			lo = mi + 1
		}

	}
	return lo*lo == num
}
