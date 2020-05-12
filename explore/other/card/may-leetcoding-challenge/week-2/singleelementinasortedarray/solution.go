package singleelementinasortedarray

// SeeAlso: https://leetcode.com/problems/single-element-in-a-sorted-array/discuss/627786/C++-O(log-n)-time-O(1)-space-or-Simple-and-clean-or-Use-xor-to-keep-track-of-odd-even-pair
func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mi := (lo + hi) / 2

		if nums[mi] == nums[mi^1] {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return nums[lo]
}
