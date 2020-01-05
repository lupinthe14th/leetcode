package main

func minInsertions(s string) int {
	r := reverse(s)
	return len(s) - lcs(s, r)
}

func reverse(s string) string {
	var ans string
	l := len(s)
	for i := 0; i < l; i++ {
		ans += string(s[l-1-i])
	}
	return ans
}

func lcs(s, t string) int {
	n, m := len(s), len(t)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[n][m]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {}
