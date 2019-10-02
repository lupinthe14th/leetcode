package removeelement

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return len(nums)
	}

	i := 0
	for _, n := range nums {
		if n == val {
			continue
		}
		nums[i] = n
		i++
	}
	nums = nums[:i]
	return i
}
