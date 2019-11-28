package intersectionoftwoarrays

import (
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	// sort
	sort.Ints(nums1)
	sort.Ints(nums2)

	// binary search
	l1, l2 := biggerSlice(nums1, nums2)

	for i := 0; i < len(l1); i++ {
		if i != 0 && l1[i] == l1[i-1] {
			continue
		}
		lo, hi := 0, len(l2)
		for lo <= hi {
			mi := (lo + hi) / 2
			if mi >= len(l2) {
				break
			}

			if l1[i] == l2[mi] {
				ans = append(ans, l1[i])
				break
			} else if l1[i] < l2[mi] {
				hi = mi - 1
			} else if l1[i] > l2[mi] {
				lo = mi + 1
			}
		}
	}
	return ans
}

func biggerSlice(a, b []int) (loS, hiS []int) {
	la, lb := len(a), len(b)

	if la < lb {
		return a, b
	} else if la > lb {
		return b, a
	} else {
		return a, b
	}
}
