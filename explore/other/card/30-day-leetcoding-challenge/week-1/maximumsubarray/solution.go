package maximumsubarray

func maxSubArray(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = nums[0]
	maxSum := dp[0]
	for i := 1; i < l; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
