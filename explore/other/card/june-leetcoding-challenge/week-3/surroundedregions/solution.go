package surroundedregions

// https://leetcode.com/problems/surrounded-regions/discuss/41661/Golang-concise-solution
func solve(board [][]byte) {
	rows := len(board)

	if rows == 0 {
		return
	}
	cols := len(board[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'O' && (row == 0 || col == 0 || row == rows-1 || col == cols-1) {
				fill(board, row, col, rows, cols)
			}
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if ch := board[row][col]; ch == 'S' {
				board[row][col] = 'O'
			} else if ch == 'O' {
				board[row][col] = 'X'
			}
		}
	}
}

func fill(board [][]byte, row, col, rows, cols int) {
	board[row][col] = 'S'
	if row-1 >= 0 && board[row-1][col] == 'O' {
		fill(board, row-1, col, rows, cols)
	}
	if col-1 >= 0 && board[row][col-1] == 'O' {
		fill(board, row, col-1, rows, cols)
	}
	if row+1 <= rows-1 && board[row+1][col] == 'O' {
		fill(board, row+1, col, rows, cols)
	}
	if col+1 <= cols-1 && board[row][col+1] == 'O' {
		fill(board, row, col+1, rows, cols)
	}
}
