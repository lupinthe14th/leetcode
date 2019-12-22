package findnumberswithevennumberofdigits

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{12, 345, 2, 6, 7896}, want: 2},
	{id: 2, input: []int{555, 901, 482, 1771}, want: 1},
}

func TestFindNumbers(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := findNumbers(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func TestNumberOfDigits(t *testing.T) {
	var tests = []struct {
		id    int
		input int
		want  int
	}{
		{id: 1, input: 0, want: 1},
		{id: 2, input: 10, want: 2},
		{id: 3, input: 100, want: 3},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := numberOfDigits(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
