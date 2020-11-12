package validsquare

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {

	dist := func(p1, p2 []int) int {
		return (p2[1]-p1[1])*(p2[1]-p1[1]) + (p2[0]-p1[0])*(p2[0]-p1[0])
	}
	check := func(p1, p2, p3, p4 []int) bool {
		return dist(p1, p2) > 0 && dist(p1, p2) == dist(p2, p3) && dist(p2, p3) == dist(p3, p4) && dist(p3, p4) == dist(p4, p1) && dist(p1, p3) == dist(p2, p4)
	}
	swap := func(p [][]int, x, y int) {
		tmp := p[x]
		p[x] = p[y]
		p[y] = tmp
	}
	var checkAllPerm func(p [][]int, l int) bool
	checkAllPerm = func(p [][]int, l int) bool {
		if l == 4 {
			return check(p[0], p[1], p[2], p[3])
		}
		out := false
		for i := l; i < 4; i++ {
			swap(p, l, i)
			out = out || checkAllPerm(p, l+1)
			swap(p, l, i)
		}
		return out
	}
	p := [][]int{p1, p2, p3, p4}
	return checkAllPerm(p, 0)
}
