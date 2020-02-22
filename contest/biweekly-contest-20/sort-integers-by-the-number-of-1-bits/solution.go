package sortintegersbythenumberof1bits

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {

	nums := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		nums[i] = []int{bitCount(fmt.Sprintf("%b", arr[i])), arr[i]}
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i][0] == nums[j][0] {
			return nums[i][1] < nums[j][1]
		}
		return nums[i][0] < nums[j][0]
	})

	ans := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ans[i] = nums[i][1]
	}
	return ans
}

func bitCount(s string) int {
	var c int
	for _, r := range s {
		if r == '1' {
			c++
		}
	}
	return c
}
