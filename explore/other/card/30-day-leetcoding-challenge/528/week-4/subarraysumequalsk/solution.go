package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	c := 0
	l := len(nums)
	sumA := make([]int, l+1)
	sumA[0] = 0
	for i := 1; i < l+1; i++ {
		sumA[i] = nums[i-1] + sumA[i-1]
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l+1; j++ {
			if sumA[j]-sumA[i] == k {
				c++
			}
		}
	}
	return c
}
