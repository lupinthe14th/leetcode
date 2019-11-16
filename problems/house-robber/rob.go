package houserobber

import "log"

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, 2)

	for j := 0; j < 2; j++ {
		log.Printf("j: %d", j)
		i := j
		dp[i] = nums[i]
		for i < n-2 {
			log.Printf("i: %d", i)
			dp[j] = dp[j] + nums[i+2]
			i = i + 2
		}
	}
	log.Printf("dp: %v", dp)
	maxMoney := max(dp[0], dp[1])
	return maxMoney
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
