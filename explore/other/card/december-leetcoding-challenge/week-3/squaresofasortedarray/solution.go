package squaresofasortedarray

func sortedSquares(nums []int) []int {
	out := make([]int, len(nums))
	k := len(nums) - 1
	s, e := 0, len(nums)-1

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for !(s > e) {
		if abs(nums[s]) > abs(nums[e]) {
			out[k] = nums[s] * nums[s]
			s++
			k--
		} else {
			out[k] = nums[e] * nums[e]
			e--
			k--
		}
	}
	return out
}
