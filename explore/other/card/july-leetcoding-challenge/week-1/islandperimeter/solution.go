package islandperimeter

func islandPerimeter(grid [][]int) int {
	out := 0
	r, c := len(grid), len(grid[0])

	var helper func(i, j int)

	helper = func(i, j int) {
		if i >= r || j >= c || i < 0 || j < 0 {
			return
		}
		if grid[i][j] == 1 {
			out--
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				out += 4
				helper(i+1, j)
				helper(i-1, j)
				helper(i, j+1)
				helper(i, j-1)
			}
		}
	}
	return out
}
