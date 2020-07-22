package squaresofasortedarray

func sortedSquares(A []int) []int {
	out := make([]int, len(A))
	for l, r, i := 0, len(A)-1, len(A)-1; l <= r; i-- {
		ll := A[l] * A[l]
		rr := A[r] * A[r]
		if ll <= rr {
			out[i] = rr
			r--
		} else {
			out[i] = ll
			l++
		}
	}
	return out
}
