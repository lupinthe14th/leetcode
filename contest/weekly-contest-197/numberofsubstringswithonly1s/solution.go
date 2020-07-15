package numberofsubstringswithonly1s

func numSub(s string) int {
	const MOD = 1e9 + 7
	var memo [1e5 + 1]int
	out := 0

	hi, c := 0, 0
	for i := range s {
		if s[i] == '0' {
			c = 0
		} else {
			c++
		}
		hi = max(hi, c)
		memo[c]++
	}
	if memo[0] == len(s) {
		return 0
	}

	for i := 1; i <= hi; i++ {
		out += i * memo[i] % MOD
	}
	return out % MOD
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
