package besttimetobuyandsellstockwithcooldown

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/discuss/75928/Share-my-DP-solution-(By-State-Machine-Thinking)
func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	s0 := make([]int, l)
	s1 := make([]int, l)
	s2 := make([]int, l)
	s1[0] = -prices[0]
	s0[0] = 0
	s2[0] = -1 << 31
	for i := 1; i < l; i++ {
		s0[i] = max(s0[i-1], s2[i-1])
		s1[i] = max(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
	}
	return max(s0[l-1], s2[l-1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
