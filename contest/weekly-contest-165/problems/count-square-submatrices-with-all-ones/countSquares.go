package countsquaresubmatriceswithallones

// refarence: https://leetcode.com/problems/count-square-submatrices-with-all-ones/discuss/441399/C%2B%2B-DP-Solution-with-VIDEO-explenation.-(By.-Tushar-Roy)
func countSquares(matrix [][]int) int {
	var ans int
	R, C := len(matrix), len(matrix[0])

	DP := make([][]int, R)
	for i := range DP {
		DP[i] = make([]int, C)
	}

	for i := 0; i < R; i++ {
		DP[i][0] = matrix[i][0]
	}

	for j := 0; j < C; j++ {
		DP[0][j] = matrix[0][j]
	}

	for i := 1; i < R; i++ {
		for j := 1; j < C; j++ {
			if matrix[i][j] == 1 {
				DP[i][j] = min(DP[i-1][j-1], min(DP[i][j-1], DP[i-1][j])) + 1
			}
		}
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			ans += DP[i][j]
		}
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
