package countsquaresubmatriceswithallones

func countSquares(matrix [][]int) int {
	out := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] > 0 && i > 0 && j > 0 {
				matrix[i][j] = min(min(matrix[i-1][j-1], matrix[i][j-1]), matrix[i-1][j]) + 1
			}
			out += matrix[i][j]
		}
	}
	return out
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
