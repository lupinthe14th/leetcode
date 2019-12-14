package removecoveredintervals

func removeCoveredIntervals(intervals [][]int) int {
	c := len(intervals)
	if c == 1 {
		return c
	}
	for i := 0; i <= len(intervals); i++ {
		for j := i + 1; j <= len(intervals)-1; j++ {
			if intervals[j][0] <= intervals[i][0] && intervals[i][1] <= intervals[j][1] {
				c--
			}
		}
	}
	return c
}
