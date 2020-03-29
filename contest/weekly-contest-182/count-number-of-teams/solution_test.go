package countnumberofteams

import (
	"fmt"
	"testing"
)

type Case struct {
	in   []int
	want int
}

var cases = []Case{
	{in: []int{2, 5, 3, 4, 1}, want: 3},
	{in: []int{2, 1, 3}, want: 0},
	{in: []int{1, 2, 3, 4}, want: 4},
}

func TestNumTeams(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := numTeams(c.in)
			if got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
