package arrangingcoins

func arrangeCoins(n int) int {
	lo, hi := 0, n
	for lo <= hi {
		mi := (lo + hi) / 2
		cur := mi * (mi + 1) / 2
		if cur == n {
			return mi
		}
		if cur > n {
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}
	return hi
}
