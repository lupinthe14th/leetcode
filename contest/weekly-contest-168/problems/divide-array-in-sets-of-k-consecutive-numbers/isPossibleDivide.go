package dividearrayinsetsofkconsecutivenumbers

import (
	"sort"
)

func isPossibleDivide(nums []int, k int) bool {
	l := len(nums)

	if l%k != 0 {
		return false
	}
	if k == 1 {
		return true
	}

	sort.Ints(nums)
	m := make(map[int]int, l)
	for _, v := range nums {
		m[v]++
	}

	for i := 0; i < l; i++ {
		var lo int
		if m[nums[i]] > 0 {
			lo = nums[i]
			m[lo]--
		} else {
			continue
		}
		for j := 0; j < k-1; j++ {
			n := lo + 1
			if m[n] <= 0 {
				return false
			}
			m[n]--
			lo = n
		}
	}
	return true
}
