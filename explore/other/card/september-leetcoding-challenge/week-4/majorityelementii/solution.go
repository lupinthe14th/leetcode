package majorityelementii

func majorityElement(nums []int) []int {
	count1, count2 := 0, 0
	candidate1, candidate2 := 0, 1
	for _, n := range nums {
		if candidate1 == n {
			count1++
		} else if candidate2 == n {
			count2++
		} else if count1 == 0 {
			candidate1 = n
			count1 = 1
		} else if count2 == 0 {
			candidate2 = n
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	out := make([]int, 0, 2)
	count1, count2 = 0, 0

	for _, n := range nums {
		if candidate1 == n {
			count1++
		}
		if candidate2 == n {
			count2++
		}
	}

	n := len(nums)

	if count1 > n/3 {
		out = append(out, candidate1)
	}
	if count2 > n/3 {
		out = append(out, candidate2)
	}
	return out
}
