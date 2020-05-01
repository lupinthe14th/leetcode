package maximalsquare

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}
	maxSquareLen := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				sqlen := 1
				flag := true
				for sqlen+i < m && sqlen+j < n && flag {
					for k := j; k <= sqlen+j; k++ {
						if matrix[i+sqlen][k] == '0' {
							flag = false
							break
						}
					}
					for k := i; k <= sqlen+i; k++ {
						if matrix[k][j+sqlen] == '0' {
							flag = false
							break
						}

					}
					if flag {
						sqlen++
					}
				}
				maxSquareLen = max(maxSquareLen, sqlen)
			}

		}
	}
	return maxSquareLen * maxSquareLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
