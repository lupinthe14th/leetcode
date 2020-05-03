package minimumpathsum

// SeeAlso:
// - https://leetcode.com/problems/minimum-path-sum/discuss/23457/C++-DP
//
// This is a typical DP problem. Suppose the minimum path sum of arriving at
// point (i, j) is s[i][j], then the state equation is s[i][j] =
// min(s[i-1[j],s[i][j-1]) + grid[i][j].
//
// Well, some boundary conditions need to be handled. The boundary conditions
// happen on the topmost row (s[i-1][j] does not exit) and the leftmost column
// (s[i][j-1] does not exist). Suppose grid is like [1, 1, 1, 1], then the
// minimum sum to arrive at each point is simply an accumulation of previous
// points and the result is [1, 2, 3, 4].
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	sum := make([][]int, m)
	for i := range sum {
		sum[i] = make([]int, n)
	}
	sum[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		sum[0][j] = sum[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sum[i][j] = min(sum[i-1][j], sum[i][j-1]) + grid[i][j]
		}
	}
	return sum[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
