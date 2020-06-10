package searchinsertposition

func searchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
