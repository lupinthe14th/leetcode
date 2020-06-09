package issubsequence

func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n] == m
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
