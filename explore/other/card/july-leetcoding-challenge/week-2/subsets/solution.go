package subsets

// Approach 3: Lexicographic (Binary Sorted) Subsets
func subsets(nums []int) [][]int {
	out := make([][]int, 0)
	for s := 0; s < 1<<uint(len(nums)); s++ {
		tmp := make([]int, 0)
		for i := 0; i < len(nums); i++ {
			if s>>uint(i)&1 == 1 {
				tmp = append(tmp, nums[i])
			}
		}
		out = append(out, tmp)
	}
	return out
}
