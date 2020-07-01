package perfectsquares

// https://leetcode.com/problems/perfect-squares/discuss/71488/Summary-of-4-different-solutions-(BFS-DP-static-DP-and-mathematics)
func numSquares(n int) int {
	if n <= 0 {
		return 0
	}

	// dp[i] = the least number of perfect square numbers
	// whitch sum to i. Note that dp[0] is 0
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1 << 31
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		// for each i, it must be the sum of some number (i - j*j) and
		// a perfect square number (j*j).
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
