package flippinganimage

func flipAndInvertImage(A [][]int) [][]int {

	c := len(A[0])
	invert := func(a int) int {
		if a == 1 {
			return 0
		}
		return 1
	}
	for i := range A {
		for j := 0; j < c/2; j++ {
			A[i][j], A[i][c-j-1] = A[i][c-j-1], A[i][j]
		}
	}

	for i := range A {
		for j := range A[i] {
			A[i][j] = invert(A[i][j])
		}
	}
	return A
}
