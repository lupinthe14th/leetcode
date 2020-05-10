package counttripletsthatcanformtwoarraysofequalxor

func countTriplets(arr []int) int {
	c := 0
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j; k < l; k++ {
				a := helper(arr, i, j-1)
				b := helper(arr, j, k)
				if a == b {
					c++
				}
			}
		}
	}
	return c
}

func helper(arr []int, s, e int) int {
	ans := arr[s]
	for i := s + 1; i <= e; i++ {
		ans ^= arr[i]
	}
	return ans
}
