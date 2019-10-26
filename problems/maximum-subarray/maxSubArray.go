package maximumsubarray

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	maxSum := dp[0]

	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}

		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
