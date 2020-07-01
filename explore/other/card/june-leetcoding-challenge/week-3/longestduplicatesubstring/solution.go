package longestduplicatesubstring

// https://leetcode.com/problems/longest-duplicate-substring/discuss/290871/Python-Binary-Search
// https://leetcode.com/problems/longest-duplicate-substring/discuss/291048/C++-solution-using-Rabin-Karp-and-binary-search-with-detailed-explaination
func longestDupSubstring(S string) string {
	const MOD = 1<<64 - 1
	memo := make([]uint64, len(S))
	for i, r := range S {
		memo[i] = uint64(r) - uint64('a')
	}
	p := make([]uint64, len(S))
	p[0] = 1
	for i := 1; i < len(S); i++ {
		p[i] = (p[i-1] * 26) % MOD
	}

	var helper func(l int) int

	helper = func(l int) int {
		var cur uint64 = 0
		for i := range memo[:l] {
			cur = (cur*26 + memo[i]) % MOD
		}
		seen := make(map[uint64]bool)
		seen[cur] = true
		for i := l; i < len(S); i++ {
			// log.Print(seen)
			cur = (cur*26 + memo[i] - memo[i-l]*p[l]) % MOD
			if _, ok := seen[cur]; ok {
				return i - l + 1
			}
			seen[cur] = true
		}
		return 0
	}

	ans, lo, hi := 0, 0, len(S)
	for lo < hi {
		mi := (lo + hi + 1) / 2
		pos := helper(mi)
		if pos != 0 {
			lo = mi
			ans = pos
		} else {
			hi = mi - 1
		}
	}
	return S[ans : ans+lo]
}
