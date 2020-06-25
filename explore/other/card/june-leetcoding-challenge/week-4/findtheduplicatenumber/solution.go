package findtheduplicatenumber

func findDuplicate(nums []int) int {
	l := len(nums)
	out := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] {
				out = nums[i]
				break
			}
		}
	}
	return out
}
