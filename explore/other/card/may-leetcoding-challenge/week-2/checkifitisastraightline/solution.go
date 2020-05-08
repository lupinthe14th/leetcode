package checkifitisastraightline

func checkStraightLine(coordinates [][]int) bool {
	l := len(coordinates)

	pl := coordinates[l-1]
	p0 := coordinates[0]

	a, b := 0, 0
	if pl[0]-p0[0] != 0 {
		a = (pl[1] - p0[1]) / (pl[0] - p0[0])
		b = p0[1] - a*p0[0]
	} else {
		a = 0
		b = p0[1]
	}

	for i := 1; i < l-1; i++ {
		if a*coordinates[i][0]+b != coordinates[i][1] {
			return false
		}
	}
	return true
}
