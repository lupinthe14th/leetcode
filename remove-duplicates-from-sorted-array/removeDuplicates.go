package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	seen := make(map[int]bool)
	for i, n := range nums {
		if i == 0 {
			nums = nums[:0]
		}
		if !seen[n] {
			seen[n] = true
			nums = append(nums, n)
		}
	}
	nums = nums[:len(nums)]
	return len(nums)
}
