package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)
	ans := make([]int, 2)
	var mi int
	if numbers[hi-1] <= target || numbers[lo] <= target {
		mi = hi
	} else {
		for lo < hi {
			mi = (lo + hi) / 2
			if numbers[mi] < target && numbers[mi+1] > target {
				break
			} else if numbers[mi] < target {
				lo = mi
			} else if numbers[mi] > target {
				hi = mi
			}
		}
	}
	for i := 0; i < len(numbers[:mi])+1; i++ {
		for j := i + 1; j < len(numbers[:mi]); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}

			}
		}
	}
	return ans
}
