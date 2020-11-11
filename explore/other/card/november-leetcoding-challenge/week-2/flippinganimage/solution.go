package flippinganimage

func flipAndInvertImage(A [][]int) [][]int {

	c := len(A[0])
	for i := range A {
		for j := 0; j < (c+1)/2; j++ {
			A[i][j], A[i][c-j-1] = A[i][c-j-1]^1, A[i][j]^1
		}
	}
	return A
}
