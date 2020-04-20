package numberofislands

// SeeAlso
// - https://leetcode.com/problems/number-of-islands/discuss/56340/Python-Simple-DFS-Solution
func numIslands(grid [][]byte) int {
	l := len(grid)
	if l == 0 {
		return 0
	}

	c := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				c++
			}
		}
	}
	return c
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '#'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
