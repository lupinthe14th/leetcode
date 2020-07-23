package singlenumberiii

func singleNumber(nums []int) []int {
	memo := make(map[int]int)
	for _, n := range nums {
		memo[n]++
	}

	out := make([]int, 0, len(nums))
	for k, v := range memo {
		if v == 1 {
			out = append(out, k)
		}
	}
	return out
}
