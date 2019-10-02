package searchinsertposition

func searchInsert(nums []int, target int) int {

	for i, n := range nums {
		if target < n {
			if i == 0 {
				return 0
			}
			return i

		}
		if target == n {
			return i
		}
	}
	return len(nums)
}
