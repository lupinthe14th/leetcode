package summaryranges

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	out := make([]string, 0, len(nums))
	if len(nums) == 0 {
		return out
	}
	if len(nums) == 1 {
		out = append(out, fmt.Sprint(nums[0]))
		return out
	}

	tmp := [][]int{{nums[0]}}
	t := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i]-1 {
			tmp[t] = append(tmp[t], nums[i-1])
			t++
			tmp = append(tmp, []int{nums[i]})
		} else if i == len(nums)-1 {
			tmp[t] = append(tmp[t], nums[i])
		}
	}

	for i := range tmp {
		if len(tmp[i]) == 1 {
			out = append(out, fmt.Sprint(tmp[i][0]))
		} else if tmp[i][0] == tmp[i][1] {
			out = append(out, fmt.Sprint(tmp[i][0]))
		} else {
			out = append(out, fmt.Sprintf("%v->%v", tmp[i][0], tmp[i][1]))
		}
	}
	return out
}
