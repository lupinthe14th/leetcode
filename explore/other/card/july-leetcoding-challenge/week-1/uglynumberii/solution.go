package uglynumberii

func nthUglyNumber(n int) int {
	u := make([]int, 0, n)
	var i, j, k int
	u = append(u, 1)
	for len(u) < n {
		t := min(u[i]*2, min(u[j]*3, u[k]*5))
		u = append(u, t)
		if t == u[i]*2 {
			i++
		}
		if t == u[j]*3 {
			j++
		}
		if t == u[k]*5 {
			k++
		}
	}
	return u[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
