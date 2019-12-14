package removecoveredintervals

func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 1 {
		return 1
	}
	for i := 0; i < len(intervals); i++ {
		for j := 0; j < len(intervals); j++ {
			if i == j {
				continue
			}
			if intervals[j][0] <= intervals[i][0] && intervals[i][1] <= intervals[j][1] {
				intervals = append(intervals[:i], intervals[i+1:]...)
			}
		}
	}
	return len(intervals)
}
