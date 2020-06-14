package findthelongestsubstringcontainingvowelsinevencounts

import "strings"

// https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/discuss/531840/JavaC++Python-One-Pass
func findTheLongestSubstring(s string) int {
	res, cur, n := 0, 0, len(s)
	seen := make(map[int]int)
	seen[0] = -1
	for i := 0; i < n; i++ {
		cur ^= 1<<uint(strings.IndexByte("aeiou", s[i])) + 1>>1
		if _, ok := seen[cur]; !ok {
			seen[cur] = i
		}
		res = max(res, i-seen[cur])

	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
