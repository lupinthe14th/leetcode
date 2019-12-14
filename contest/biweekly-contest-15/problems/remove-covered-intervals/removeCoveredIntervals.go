package removecoveredintervals

func removeCoveredIntervals(intervals [][]int) int {
	ans := 0
	for i := 0; i < len(intervals); i++ {
		c := 0
		for j := 0; j < len(intervals); j++ {
			if intervals[j][0] <= intervals[i][0] && intervals[i][1] <= intervals[j][1] {
				c++
			}
		}
		if c == 1 {
			ans++
		}
	}
	return ans
}
