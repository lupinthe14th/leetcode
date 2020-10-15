package houserobberii

func rob(nums []int) int {

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	helper := func(s, e int) int {
		prevTwo, prevOne, out := 0, 0, 0

		for i := s; i < e; i++ {
			out = max(prevTwo+nums[i], prevOne)
			prevTwo = prevOne
			prevOne = out

		}
		return out
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(helper(0, len(nums)-1), helper(1, len(nums)))
}
