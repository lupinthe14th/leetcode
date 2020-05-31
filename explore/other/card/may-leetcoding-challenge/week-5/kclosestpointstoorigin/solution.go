package kclosestpointstoorigin

import (
	"sort"
)

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		xi := points[i][0]
		xj := points[j][0]
		yi := points[i][1]
		yj := points[j][1]
		if xi*xi+yi*yi < xj*xj+yj*yj {
			return true
		}
		return false

	})
	return points[:K]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
