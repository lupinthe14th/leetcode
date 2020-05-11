package floodfill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var dfs func(int, int, int)
	dfs = func(row, col, color int) {
		if image[row][col] == color {
			image[row][col] = newColor
			if row-1 >= 0 {
				dfs(row-1, col, color)
			}
			if col-1 >= 0 {
				dfs(row, col-1, color)
			}
			if row+1 < len(image) {
				dfs(row+1, col, color)
			}
			if col+1 < len(image[0]) {
				dfs(row, col+1, color)
			}
		}
	}

	color := image[sr][sc]
	if newColor != color {
		dfs(sr, sc, color)
	}
	return image
}
