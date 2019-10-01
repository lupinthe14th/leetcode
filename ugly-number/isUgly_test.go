package isugly

import (
	"fmt"
	"math"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  bool
}{
	{id: 1, input: 6, want: true},
	{id: 2, input: 8, want: true},
	{id: 3, input: 14, want: false},
	{id: 4, input: int(math.Pow(-2, 31)), want: false},
	{id: 5, input: int(math.Pow(2, 31) - 1), want: false},
	{id: 6, input: 0, want: false},
	{id: 7, input: 1, want: true},
	{id: 8, input: -1, want: false},
}

func TestIsUgly(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := isUgly(tt.input)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}

}
