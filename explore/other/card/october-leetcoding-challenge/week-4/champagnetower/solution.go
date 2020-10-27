package champagnetower

import "math"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	var memo [102][102]float64
	memo[0][0] = float64(poured)
	for r := 0; r <= query_row; r++ {
		for c := 0; c <= r; c++ {
			var q float64 = (memo[r][c] - 1.0) / 2.0
			if q > 0 {
				memo[r+1][c] += q
				memo[r+1][c+1] += q
			}
		}
	}
	return math.Min(1, memo[query_row][query_glass])
}
