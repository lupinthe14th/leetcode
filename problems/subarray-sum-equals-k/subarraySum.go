package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	var c int
	l := len(nums)

	for i := 0; i < l; i++ {
		sum := nums[i]
		if sum == k {
			c++
		}
		for j := i + 1; j < l; j++ {
			sum += nums[j]
			if sum == k {
				c++
			}
		}
	}
	return c
}
