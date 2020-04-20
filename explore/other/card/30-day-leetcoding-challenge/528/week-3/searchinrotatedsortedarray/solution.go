package searchinrotatedsortedarray

// SeeAlso:
// https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14425/Concise-O(log-N)-Binary-search-solution
func search(nums []int, target int) int {
	l := len(nums)
	lo, hi := 0, l-1
	// find the index of the smallest value using binary search.
	// Loop will terminate since mid < hi, and lo or hi wil shrink by at least 1.
	// Proof by contradiction that mid < hi: if mid==hi, then lo==hi and loop would have been terminated.
	for lo < hi {
		mid := (lo + hi) / 2

		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	// lo==hi is the index of the smallest value and also the number of places rotated.
	rot := lo
	lo, hi = 0, l-1
	// The usual binary search and accounting for rotation.
	for lo <= hi {

		mid := (lo + hi) / 2
		realmid := (mid + rot) % l
		if nums[realmid] == target {
			return realmid
		}

		if nums[realmid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}
