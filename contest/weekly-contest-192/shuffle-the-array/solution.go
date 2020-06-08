package shufflethearray

func shuffle(nums []int, n int) []int {
	out := make([]int, 2*n)
	for i := 0; i < n; i++ {
		out[2*i] = nums[i]
		out[2*i+1] = nums[n+i]
	}
	return out
}
