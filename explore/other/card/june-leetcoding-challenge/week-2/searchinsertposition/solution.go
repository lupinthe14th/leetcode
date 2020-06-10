package searchinsertposition

// https://leetcode.com/explore/featured/card/june-leetcoding-challenge/540/week-2-june-8th-june-14th/3356/discuss/15101/C++-O(logn)-Binary-Search-that-handles-duplicate
func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mi := (lo + hi) >> 1

		if nums[mi] == target {
			return mi
		} else if nums[mi] < target {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return lo
}
