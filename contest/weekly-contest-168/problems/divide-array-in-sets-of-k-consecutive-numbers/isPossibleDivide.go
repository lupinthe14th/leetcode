package dividearrayinsetsofkconsecutivenumbers

import (
	"sort"
)

func isPossibleDivide(nums []int, k int) bool {
	sort.Ints(nums)
	l := len(nums)

	var sum int
	for i := 0; i < l; i++ {
		sum = nums[i]
		if sum == k {
			return true
		}
		for j := i + 1; j < l; j++ {
			sum += nums[j]
			if sum == k {
				return true
			} else if k != 0 && sum%k == 0 {
				return true
			}
		}
	}
	return false
}
