package issubsequence

// SeeAlso:
// - https://leetcode.com/problems/is-subsequence/discuss/87302/Binary-search-solution-for-follow-up-with-detailed-comments
// - https://leetcode.com/problems/is-subsequence/discuss/87292/Swift-solution-Binary-Search
func isSubsequence(s string, t string) bool {
	memo := make(map[byte][]int, 26)
	for i := range t {
		memo[t[i]] = append(memo[t[i]], i)
	}

	prev := 0
	for i := range s {
		a, ok := memo[s[i]]
		if !ok {
			return false
		}
		j := bs(prev, a)
		if j == len(a) {
			return false
		}
		prev = a[j] + 1
	}
	return true
}

func bs(n int, a []int) int {
	lo, hi := 0, len(a)-1

	for lo <= hi {
		mi := (lo + hi) >> 1
		if a[mi] < n {
			lo = mi + 1
		} else {
			hi = mi - 1
		}

	}
	return lo
}
