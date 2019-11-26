package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1
	ans := make([]int, 2)
	for lo < hi {
		tmp := numbers[lo] + numbers[hi]
		if tmp == target {
			ans[0] = lo + 1
			ans[1] = hi + 1
			return ans
		} else if tmp < target {
			lo++
		} else {
			hi--
		}
	}
	return ans
}
