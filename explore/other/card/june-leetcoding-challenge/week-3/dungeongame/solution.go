package dungeongame

// https://leetcode.com/problems/dungeon-game/discuss/330370/Go-4ms-dynamic-programming-submission-(beating-100)
func calculateMinimumHP(dungeon [][]int) int {
	rows, cols := len(dungeon)-1, len(dungeon[0])-1

	for row := rows; row >= 0; row-- {
		for col := cols; col >= 0; col-- {
			if row == rows && col == cols { // bottom-right room
				dungeon[rows][cols] = max(1, 1-dungeon[rows][cols])
				continue
			}
			if row == rows { // bottom-margin room
				dungeon[row][col] = max(1, dungeon[row][col+1]-dungeon[row][col])
				continue
			}
			if col == cols { // right-margin room
				dungeon[row][col] = max(1, dungeon[row+1][col]-dungeon[row][col])
				continue
			}
			// any other room
			dungeon[row][col] = max(1, min(dungeon[row+1][col], dungeon[row][col+1])-dungeon[row][col])

		}
	}
	return dungeon[0][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
