package anglebetweenhandsofaclock

import (
	"math"
)

func angleClock(hour int, minutes int) float64 {
	h := 30.0*float64(hour) + 0.5*float64(minutes)
	m := 6.0 * float64(minutes)

	if h >= 360 {
		h -= 360
	}
	out := math.Abs(h - m)
	if out > 180 {
		return 360 - out
	}
	return out
}
