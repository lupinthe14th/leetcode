package continuoussubarraysum

// Brute force solution
func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	var sum int
	for i := 0; i < l; i++ {
		sum = nums[i]
		for j := i + 1; j < l; j++ {
			sum += nums[j]
			if k != 0 && sum%k == 0 {
				return true
			}
		}
	}
	return false
}
