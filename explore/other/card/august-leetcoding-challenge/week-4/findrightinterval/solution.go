package findrightinterval

func findRightInterval(intervals [][]int) []int {
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

	l := len(intervals)
	out := make([]int, l)
	if l == 0 {
		return out
	}
	if l == 1 {
		out[0] = -1
		return out
	}
	memo := make(map[int]int, l)
	lo, hi := 1<<31, -1<<31
	for i := range intervals {
		memo[intervals[i][0]] = i
		lo = min(lo, intervals[i][0])
		hi = max(hi, intervals[i][0])
	}

	for i := range intervals {
		l := len(intervals[i])
		cur := intervals[i][l-1]
		if cur > hi {
			out[i] = -1
			continue
		}
		if cur <= lo {
			out[i] = memo[cur]
			continue
		}
		if cur == hi {
			out[i] = memo[cur]
			continue
		}
		if lo < cur && cur < hi {
			for j := cur; j < hi; j++ {
				if _, ok := memo[j]; ok {
					out[i] = memo[j]
					break
				}
			}
		}
	}
	return out
}
