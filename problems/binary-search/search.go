package binarysearch

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mi := (hi + lo) / 2
		if nums[mi] == target {
			return mi
		} else if nums[mi] < target {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return -1
}
