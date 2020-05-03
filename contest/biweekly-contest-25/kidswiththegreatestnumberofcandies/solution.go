package kidswiththegreatestnumberofcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	l := len(candies)
	ans := make([]bool, l)
	m := -1 << 31

	for i := range candies {
		m = max(m, candies[i])
	}

	for i := range candies {
		if candies[i]+extraCandies >= m {
			ans[i] = true
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
