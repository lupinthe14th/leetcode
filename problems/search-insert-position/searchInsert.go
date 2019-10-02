package searchinsertposition

func searchInsert(nums []int, target int) int {
	for i, n := range nums {
		if target <= n {
			return i
		}
	}
	return len(nums)
}
