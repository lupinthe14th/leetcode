package numberofdaysbetweentwodates

import (
	"time"
)

const layout = "2006-01-02"

func daysBetweenDates(date1 string, date2 string) int {
	d1, _ := time.Parse(layout, date1)
	d2, _ := time.Parse(layout, date2)
	ans := d2.Sub(d1)
	h, _ := time.ParseDuration(ans.String())
	return abs(int(h.Hours() / 24))
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
