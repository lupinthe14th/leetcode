package singleelementinasortedarray

func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mi := (lo + hi) / 2

		if (mi&1 == 0 && nums[mi] == nums[mi+1]) || (mi&1 == 1 && nums[mi] == nums[mi-1]) {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return nums[lo]
}
