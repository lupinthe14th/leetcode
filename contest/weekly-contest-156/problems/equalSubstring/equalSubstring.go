package equalsubstring

import (
	"log"
	"math"
)

func equalSubstring(s string, t string, maxCost int) int {
	log.Printf("cost: %d", maxCost)
	for i := 0; i < len(s); i++ {
		log.Printf("s: %s: %d", string(s[i]), s[i])
		log.Printf("t: %s: %d", string(t[i]), t[i])
		diff := int(s[i]) - int(t[i])
		log.Printf("diff: %d", diff)
		abs := math.Abs(float64(diff))
		log.Printf("abs: %f", abs)
		if int(abs) > maxCost {
			return i
		}
		maxCost = maxCost - int(abs)
	}
	return len(s)
}
