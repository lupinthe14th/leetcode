package cellswithoddvaluesinamatrix

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)
	// initial
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	for _, index := range indices {
		for i := 0; i < m; i++ {
			row := index[0]
			matrix[row][i]++
		}

		for i := 0; i < n; i++ {
			col := index[1]
			matrix[i][col]++
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
