package arrangingcoins

func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	cum := []int{1}
	for i := 1; cum[i-1] < n; i++ {
		cum = append(cum, i+1+cum[i-1])
	}
	lo, hi := 0, len(cum)-1

	for lo <= hi {
		mi := (lo + hi) / 2
		if cum[mi] == n {
			return mi + 1
		}
		if cum[mi] > n {
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}
	return lo
}
