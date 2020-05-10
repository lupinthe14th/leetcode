package buildanarraywithstackoperations

import (
	"reflect"
)

func buildArray(target []int, n int) []string {
	ans := make([]string, 0)
	nums := make([]int, 0)
	j := 0
	for i := 0; i < n; i++ {
		ans = append(ans, "Push")
		nums = append(nums, i+1)

		if target[j] != i+1 {
			ans = append(ans, "Pop")
			nums = nums[:len(nums)-1]
			j--
		}

		if reflect.DeepEqual(target, nums) {
			break
		}
		j++

	}
	return ans
}
