package binarysearch

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mi := (hi + lo) / 2
		if nums[mi] == target {
			return mi
		} else if nums[mi] > target {
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}
	return -1
}
