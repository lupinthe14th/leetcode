package cheapestflightswithinkstops

// https://leetcode.com/problems/cheapest-flights-within-k-stops/discuss/686774/SUGGESTION-FOR-BEGINNERS-SIMPLE-STEPS-BFS-or-DIJKSHTRA-or-DP-DIAGRAM
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	const INT_MAX = 1 << 31
	// dp[i][j] = Dist to reach j using atmost edges from src
	dp := make([][]int, K+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INT_MAX
		}
	}
	for i := 0; i <= K+1; i++ {
		// Dist to reach src always zero
		dp[i][src] = 0
	}

	for i := 1; i <= K+1; i++ {
		for _, f := range flights {
			u, v, w := f[0], f[1], f[2]

			if dp[i-1][u] != INT_MAX {
				dp[i][v] = min(dp[i][v], dp[i-1][u]+w)
			}
		}
	}
	if dp[K+1][dst] == INT_MAX {
		return -1
	}
	return dp[K+1][dst]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
