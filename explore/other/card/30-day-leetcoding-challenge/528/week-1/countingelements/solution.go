package countingelements

func countElements(arr []int) int {
	seen := make(map[int]bool, len(arr))
	for _, v := range arr {
		seen[v] = true
	}
	c := 0
	for _, v := range arr {
		if seen[v+1] {
			c++
		}
	}
	return c
}
