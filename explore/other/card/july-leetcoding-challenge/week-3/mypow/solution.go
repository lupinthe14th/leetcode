package mypow

// https://leetcode.com/problems/powx-n/discuss/473575/100-percentile-CPP-solution
func myPow(x float64, n int) float64 {
	out := 1.0
	i := n

	for n != 0 {
		if n&1 == 1 {
			out *= x
		}
		x *= x
		n = n / 2
	}

	if i >= 0 {
		return out
	}
	return 1 / out
}
