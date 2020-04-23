package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	c := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		sum := 0
		for j := i; j < l; j++ {
			sum += nums[j]
			if sum == k {
				c++
			}
		}
	}
	return c
}
