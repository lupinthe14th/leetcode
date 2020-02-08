package anglebetweenhandsofaclock

func angleClock(hour int, minutes int) float64 {
	// 時針
	var h float64
	if hour == 12 {
		h = 0.5 * float64(minutes)
	} else {
		h = float64(hour)*30 + 0.5*float64(minutes)
	}
	// 分針
	var m float64 = 6 * float64(minutes)

	ans := min(abs(h, m), abs(m, h))
	if ans > 180 {
		return 360 - ans
	}
	return ans
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func abs(x, y float64) float64 {
	z := x - y
	if z < 0 {
		return z * -1
	}
	return z
}
