package robotboundedincircle

func isRobotBounded(instructions string) bool {

	type robot struct {
		p [2]int
		f int // north: 0, east: 1, south: 2, west: 3
	}

	r := robot{p: [2]int{0, 0}, f: 0}
	goStraight := func() {
		switch r.f {
		case 0:
			r.p[1]++
		case 1:
			r.p[0]++
		case 2:
			r.p[1]--
		case 3:
			r.p[0]--
		}
	}

	turn := func(d rune) {
		if d == 'L' {
			if r.f == 0 {
				r.f = 3
			} else {
				r.f--
			}
		} else {
			if r.f == 3 {
				r.f = 0
			} else {
				r.f++
			}
		}
	}

	for _, instruction := range instructions {
		if instruction == 'G' {
			goStraight()
		} else {
			turn(instruction)
		}
	}
	if r.p == [2]int{0, 0} || r.f != 0 {
		return true
	}
	return false
}
