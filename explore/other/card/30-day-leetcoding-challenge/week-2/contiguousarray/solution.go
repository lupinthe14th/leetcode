package contiguousarray

// SeeAlso: https://leetcode.com/problems/contiguous-array/discuss/577324/C++-SHORT-Code-with-Explanation
//
// Declare two variables count,length.
//
// First we traverse the elements and if it is zero count++, if 1 count--.
//
// First check if count == 0 then compare it with maximum length, if it is not.
//
// Then we check if the same count ever happend before, if yes then it there
// are equal 0&1 in between in these indices, and if not insert that count in
// map.
func findMaxLength(nums []int) int {
	c, l := 0, 0
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			c++
		} else {
			c--
		}

		if c == 0 {
			l = max(l, i+1)
		} else if mc, ok := m[c]; ok {
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
