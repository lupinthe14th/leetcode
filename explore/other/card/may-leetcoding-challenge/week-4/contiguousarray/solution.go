package contiguousarray

func findMaxLength(nums []int) int {
	c, l := 0, 0
	m := make(map[int]int, len(nums))
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			c++
		} else {
			c--
		}
		if mc, ok := m[c]; ok {
			l = max(l, i-mc)
		} else {
			m[c] = i
		}
	}
	return l
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
