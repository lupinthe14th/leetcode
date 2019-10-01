package removeduplicatesfromsortedarray

import "log"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	i := 0
	for _, v := range nums[1:] {
		if v == nums[i] {
			continue
		}
		nums[i+1] = v
		i++
	}
	log.Println(nums)
	return i + 1
}
