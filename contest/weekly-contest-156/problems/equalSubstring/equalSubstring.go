package equalsubstring

import (
	"math"
)

func equalSubstring(s string, t string, maxCost int) int {
	var counter int
	for i := 0; i < len(s); i++ {
		diff := int(s[i]) - int(t[i])
		abs := math.Abs(float64(diff))
		if int(abs) > maxCost {
			continue
		}
		counter++
		maxCost = maxCost - int(abs)
	}
	return counter
}
