package minimumcostfortickets

func mincostTickets(days []int, costs []int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	var memo [366]int
	dayset := make(map[int]bool)
	for _, day := range days {
		dayset[day] = true
	}
	var dp func(i int) int
	dp = func(i int) int {
		if i > 365 {
			return 0
		}
		if memo[i] != 0 {
			return memo[i]
		}

		out := 0
		if dayset[i] {
			out = min(dp(i+1)+costs[0], min(dp(i+7)+costs[1], dp(i+30)+costs[2]))
		} else {
			out = dp(i + 1)
		}
		memo[i] = out
		return out
	}

	return dp(1)
}
