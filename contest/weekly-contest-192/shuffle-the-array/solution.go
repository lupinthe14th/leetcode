package shufflethearray

func shuffle(nums []int, n int) []int {
	out := make([]int, len(nums))
	for i, j := 0, 0; i < len(nums)/2; i++ {
		out[j] = nums[i]
		j++
		out[j] = nums[(len(nums)/2)+i]
		j++
	}
	return out
}
