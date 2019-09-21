package maxnumberofapples

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{100, 200, 150, 1000}, want: 4},
	{id: 2, input: []int{900, 950, 800, 1000, 700, 800}, want: 5},
	{id: 3, input: []int{988, 641, 984, 635, 461, 527, 491, 610, 274, 104, 348, 468, 220, 837, 126, 111, 536, 368, 89, 201, 589, 481, 849, 708, 258, 853, 563, 868, 92, 76, 566, 398, 272, 697, 584, 851, 302, 766, 88, 428, 276, 797, 460, 244, 950, 134, 838, 161, 15, 330}, want: 23},
}

func TestMaxNumberOfApples(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := maxNumberOfApples(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
