package findluckyintegerinanarray

func findLucky(arr []int) int {
	var ans int = -1
	memo := make(map[int]int, len(arr))
	for _, v := range arr {
		memo[v]++
	}
	for _, v := range memo {
		if v == memo[v] {
			ans = max(ans, v)
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
