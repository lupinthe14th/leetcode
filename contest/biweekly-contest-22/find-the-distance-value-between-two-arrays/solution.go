package findthedistancevaluebetweentwoarrays

import "log"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var ans int
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			log.Printf("i: %d, abs: %d, arr1: %d, arr2: %d", i, abs(arr1[i], arr2[j]), arr1[i], arr2[j])
			if abs(arr1[i], arr2[j]) <= d {
				return i
			}
		}
	}
	return ans
}

func abs(x, y int) int {
	ans := x - y
	if ans < 0 {
		return ans * -1
	}
	return ans
}
