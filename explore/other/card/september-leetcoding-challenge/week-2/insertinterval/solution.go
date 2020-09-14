package insertinterval

// https://leetcode.com/problems/insert-interval/discuss/304135/Go-8ms-6.2mb
func insert(intervals [][]int, newInterval []int) [][]int {
	out := make([][]int, 0)
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i, v := range intervals {
		if v[1] < newInterval[0] {
			out = append(out, v)
			continue
		}

		if v[0] > newInterval[1] {
			out = append(out, newInterval)
			out = append(out, intervals[i:]...)
			return out
		}

		newInterval[0] = min(newInterval[0], v[0])
		newInterval[1] = max(newInterval[1], v[1])
	}
	return append(out, newInterval)
}
