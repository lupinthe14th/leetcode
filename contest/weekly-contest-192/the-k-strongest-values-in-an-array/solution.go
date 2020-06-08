package thekstrongestvaluesinanarray

import (
	"sort"
)

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	mi := (len(arr) - 1) / 2
	m := arr[mi]

	sort.Slice(arr, func(i, j int) bool {
		if abs(arr[i]-m) == abs(arr[j]-m) {
			return arr[i] > arr[j]
		}
		return abs(arr[i]-m) > abs(arr[j]-m)
	})
	return arr[:k]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
