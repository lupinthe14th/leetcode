package houserobber

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, 2)

	var maxMoney int
	if n > 1 {

		for j := 0; j < 2; j++ {
			i := j
			dp[i] = nums[i]
			for i < n-2 {
				dp[j] = dp[j] + nums[i+2]
				i = i + 2
			}
		}
		maxMoney = max(dp[0], dp[1])
	} else {
		maxMoney = nums[0]
	}
	return maxMoney
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
