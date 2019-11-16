package houserobber

import "log"

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = nums[0]

	for i := 1; i < n; i++ {
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
		log.Printf("nums: %v", nums)
		log.Printf("dp[i]: %d", dp[i])
		log.Printf("dp[i-1]: %d", dp[i-1])
		log.Printf("nums[i]: %d", nums[i])
	}
	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
