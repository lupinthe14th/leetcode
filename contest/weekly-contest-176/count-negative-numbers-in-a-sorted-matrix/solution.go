package countnegativenumbersinasortedmatrix

func countNegatives(grid [][]int) int {
	var ans int

	for _, r := range grid {
		for _, c := range r {
			if c < 0 {
				ans++
			}
		}
	}
	return ans
}
