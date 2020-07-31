package wordbreakii

import "strings"

// seealso: https://leetcode.com/problems/word-break-ii/discuss/596188/DP-solution
func wordBreak(s string, wordDict []string) []string {
	out := make([]string, 0)

	var dfs func(dp [][]string, cur string, index int)
	dfs = func(dp [][]string, cur string, index int) {
		if index < 0 {
			out = append(out, strings.TrimSpace(cur))
			return
		}

		for i := 0; i < len(dp[index]); i++ {
			nextIndex := index - len(dp[index][i])
			next := dp[index][i] + " " + cur

			dfs(dp, next, nextIndex)
		}
	}

	n := len(s)
	dp := make([][]string, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]string, 0)
	}

	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}

	for r := 0; r < n; r++ {
		for l := 0; l <= r; l++ {
			cur := s[l : r+1]
			if m[cur] {
				if l < 1 || len(dp[l-1]) > 0 {
					dp[r] = append(dp[r], cur)
				}
			}
		}
	}

	if len(dp[n-1]) == 0 {
		return out
	}

	dfs(dp, "", n-1)
	return out
}
