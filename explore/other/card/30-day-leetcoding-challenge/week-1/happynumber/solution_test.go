package happynumber

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {
	type c struct {
		in   int
		want bool
	}

	var cases = []c{
		{in: 19, want: true},
		{in: 2, want: false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isHappy(c.in)
			if got != c.want {
				t.Fatalf("got: %v, want: %v", got, c.want)
			}
		})
	}
}

func TestSumOfSquresOfDigits(t *testing.T) {
	type c struct {
		in, want int
	}

	var cases = []c{
		{in: 19, want: 82},
		{in: 82, want: 68},
		{in: 68, want: 100},
		{in: 100, want: 1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sumOfSquaresOfDigits(c.in)
			if got != c.want {
				t.Fatalf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
