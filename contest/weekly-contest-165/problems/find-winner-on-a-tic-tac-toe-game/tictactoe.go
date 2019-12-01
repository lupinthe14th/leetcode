package findwinneronatictactoegame

func tictactoe(moves [][]int) string {

	if len(moves) < 5 {
		return "Pending"
	}

	matrix := make([][]string, 3)
	for i := range matrix {
		matrix[i] = make([]string, 3)
	}
	var c int
	for i, move := range moves {
		c++
		x := move[0]
		y := move[1]
		if i%2 == 0 {
			matrix[x][y] = "A"
		} else {
			matrix[x][y] = "B"
		}

		for x := 0; x < 3; x++ {
			if matrix[x][0] != "" && matrix[x][0] == matrix[x][1] && matrix[x][1] == matrix[x][2] {
				return matrix[x][0]
			}
		}
		for y := 0; y < 3; y++ {
			if matrix[0][y] != "" && matrix[0][y] == matrix[1][y] && matrix[1][y] == matrix[2][y] {
				return matrix[0][y]
			}
		}
		if matrix[0][0] != "" && matrix[0][0] == matrix[1][1] && matrix[1][1] == matrix[2][2] {
			return matrix[0][0]
		}
		if matrix[2][0] != "" && matrix[2][0] == matrix[1][1] && matrix[1][1] == matrix[0][2] {
			return matrix[2][0]
		}

	}
	if c == 9 {
		return "Draw"
	}
	return "Pending"
}
