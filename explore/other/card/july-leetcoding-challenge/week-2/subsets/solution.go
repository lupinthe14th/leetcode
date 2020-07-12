package subsets

// Approach 1: Cascading
func subsets(nums []int) [][]int {
	out := [][]int{}
	out = append(out, []int{})

	for _, num := range nums {
		subsets := [][]int{}
		for _, curr := range out {
			curr = append([]int{num}, curr...)
			subsets = append(subsets, curr)
		}
		for _, curr := range subsets {
			out = append(out, curr)
		}
	}
	return out
}
