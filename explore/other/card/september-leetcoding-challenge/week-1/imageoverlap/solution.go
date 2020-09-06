package imageoverlap

func largestOverlap(A [][]int, B [][]int) int {
	row := len(A)
	col := len(A[0])

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	helper := func(x, y int) int {
		asc := 0
		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				if i+x >= row || j+y >= col {
					continue
				}
				if A[i][j] == 1 && A[i][j] == B[i+x][j+y] {
					asc++
				}
			}
		}
		dsc := 0
		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				if i-x < 0 || j-y < 0 {
					continue
				}
				if A[i][j] == 1 && A[i][j] == B[i-x][j-y] {
					dsc++
				}
			}
		}
		return max(asc, dsc)
	}
	out := -1 << 31
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			out = max(out, helper(i, j))
		}
	}
	return out
}
