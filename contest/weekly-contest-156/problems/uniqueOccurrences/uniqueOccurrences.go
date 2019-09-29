package uniqueoccurrences

import "log"

func uniqueOccurrences(arr []int) bool {
	var counter = make(map[int]int, 0)
	for _, v := range arr {
		counter[v]++
	}
	log.Print(counter)

	seen := make(map[int]bool)

	for _, v := range counter {
		if !seen[v] {
			seen[v] = true
		} else {
			return false
		}
	}
	return true
}
