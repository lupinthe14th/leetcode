package intervallistintersections

func intervalIntersection(A [][]int, B [][]int) [][]int {
	out := [][]int{}
	i, j := 0, 0

	for i < len(A) && j < len(B) {
		// Let's check if A[i] intersects B[j]
		// lo - the startpoint of the intersection
		// hi - the endpoint of the intersection
		lo := max(A[i][0], B[j][0])
		hi := min(A[i][1], B[j][1])
		if lo <= hi {
			out = append(out, []int{lo, hi})
		}

		// Remove the interval with the smallest endpoint
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
