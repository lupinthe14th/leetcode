package numberofdaysbetweentwodates

import (
	"fmt"
	"testing"
)

type Case struct {
	in   in
	want int
}

type in struct {
	date1, date2 string
}

var cases = []Case{
	{in: in{date1: "2019-06-29", date2: "2019-06-30"}, want: 1},
	{in: in{date1: "2020-01-15", date2: "2019-12-31"}, want: 15},
}

func TestDaysBetweenDates(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := daysBetweenDates(c.in.date1, c.in.date2)
			if got != c.want {
				t.Errorf("%d, want: %d", got, c.want)
			}
		})
	}
}
