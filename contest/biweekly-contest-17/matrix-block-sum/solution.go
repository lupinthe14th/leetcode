package main

func matrixBlockSum(mat [][]int, K int) [][]int {
	m := len(mat)
	n := len(mat[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rmax, rmin := betweenlist(i, K, m)
			cmax, cmin := betweenlist(j, K, n)
			var answer int
			for r := rmin; r <= rmax; r++ {
				for c := cmin; c <= cmax; c++ {
					answer += mat[r][c]
				}
			}

			ans[i][j] = answer
		}
	}
	return ans
}

func betweenlist(n, K, l int) (max int, min int) {
	min = n - K
	if min < 0 {
		min = 0
	}
	max = n + K
	if max > l-1 {
		max = l - 1
	}
	return max, min
}

func main() {}
