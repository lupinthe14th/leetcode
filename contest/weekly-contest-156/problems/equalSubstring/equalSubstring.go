package equalsubstring

import (
	"math"
	"sort"
)

func equalSubstring(s string, t string, maxCost int) int {
	var counter int
	stack := make([]int, 0)
	cost := maxCost
	for j := 0; j < len(s); j++ {
		counter = 0
		cost = maxCost
		for i := j; i < len(s); i++ {
			diff := int(s[i]) - int(t[i])
			abs := math.Abs(float64(diff))
			if int(abs) > cost {
				stack = append(stack, counter)
				break
			}
			counter++
			cost = cost - int(abs)
		}
		stack = append(stack, counter)
	}
	if len(stack) == 0 {
		return 0
	}
	sort.Ints(stack)
	return stack[len(stack)-1]
}
