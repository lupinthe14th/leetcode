package pathcrossing

func isPathCrossing(path string) bool {
	var memo [1000][1000]int
	var i, j int = 500, 500
	memo[i][j] = 1
	for _, v := range path {
		switch v {
		case 'N':
			i++
			if memo[i][j] == 1 {
				return true
			} else {
				memo[i][j] = 1
			}
		case 'E':
			j++
			if memo[i][j] == 1 {
				return true
			} else {
				memo[i][j] = 1
			}
		case 'S':
			i--
			if memo[i][j] == 1 {
				return true
			} else {
				memo[i][j] = 1
			}
		case 'W':
			j--
			if memo[i][j] == 1 {
				return true
			} else {
				memo[i][j] = 1
			}
		}
	}
	return false
}
