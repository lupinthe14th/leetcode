package maxnumberofapples

import "log"

func maxNumberOfApples(arr []int) int {
	var sum int
	for i, w := range arr {
		log.Print(i)
		sum = sum + w
		if sum > 5000 {
			return i
		}
	}
	return len(arr)
}
