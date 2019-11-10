package cellswithoddvaluesinamatrix

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)
	// initial
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	for _, v := range indices {
		for i := 0; i < m; i++ {
			matrix[v[0]][i] = matrix[v[0]][i] + 1
		}

		for i := 0; i < n; i++ {
			matrix[i][v[1]] = matrix[i][v[1]] + 1
		}
	}

	var c int
	for _, arr := range matrix {
		for _, num := range arr {
			if num&1 == 1 {
				c++
			}
		}
	}

	return c
}
