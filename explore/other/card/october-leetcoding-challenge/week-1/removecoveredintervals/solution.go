package removecoveredintervals

import (
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	overlap := 0
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]

	})
	for i := range intervals {
		a := intervals[i][0]
		b := intervals[i][1]
		for j := range intervals {
			if i == j {
				continue
			}
			c := intervals[j][0]
			d := intervals[j][1]

			if c <= a && b <= d {
				overlap++
				break
			}

		}
	}
	return len(intervals) - overlap
}
