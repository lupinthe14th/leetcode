package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)
	ans := make([]int, 2)

	for lo < hi {
		mi := (lo + hi) / 2
		if numbers[mi] > target && numbers[mi-1] < target {
			ans = []int{mi - 1, mi}
			break
		} else if numbers[mi] < target {
			lo = mi
		} else if numbers[mi] > target {
			hi = mi
		}
	}
	return ans
}
