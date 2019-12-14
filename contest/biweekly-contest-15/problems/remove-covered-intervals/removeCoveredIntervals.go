package removecoveredintervals

func removeCoveredIntervals(intervals [][]int) int {
	c := len(intervals)
	if c == 1 {
		return c
	}
	for i := 0; i < len(intervals); i++ {
		for j := 0; j < len(intervals); j++ {
			if i == j {
				continue
			}
			if intervals[j][0] <= intervals[i][0] && intervals[i][1] <= intervals[j][1] {
				c--
			}
		}
	}
	return c
}
