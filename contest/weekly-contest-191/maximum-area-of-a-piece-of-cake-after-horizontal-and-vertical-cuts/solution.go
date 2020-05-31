package maximumareaofapieceofcakeafterhorizontalandverticalcuts

import (
	"log"
	"sort"
)

const MOD = 1e9 + 7

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	hih := horizontalCuts[0]
	for i := 0; i < len(horizontalCuts); i++ {
		if i == len(horizontalCuts)-1 {
			hih = max(hih, h-horizontalCuts[i])
		} else {
			hih = max(hih, horizontalCuts[i+1]-horizontalCuts[i])
		}
	}
	sort.Ints(verticalCuts)
	hiw := verticalCuts[0]
	for i := 0; i < len(verticalCuts); i++ {
		log.Print(verticalCuts[i])
		if i == len(verticalCuts)-1 {
			hiw = max(hiw, w-verticalCuts[i])
		} else {
			hiw = max(hiw, verticalCuts[i+1]-verticalCuts[i])
		}
	}
	return (hih * hiw) % MOD
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
