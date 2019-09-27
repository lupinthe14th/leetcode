package nthuglynumber

import (
	"fmt"
	"testing"
)

func TestLCM(t *testing.T) {
	var cases = []struct {
		id      int
		a, b, c int
		want    int
	}{
		{id: 1, a: 2, b: 3, c: 5, want: 30},
		{id: 2, a: 2, b: 3, c: 4, want: 24},
		{id: 3, a: 2, b: 11, c: 13, want: 286},
	}
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := lcm(tt.a, tt.b, tt.c)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func TestNthUglyNumber(t *testing.T) {
	var cases = []struct {
		id         int
		n, a, b, c int
		want       int
	}{
		{id: 1, n: 3, a: 2, b: 3, c: 5, want: 4},
		{id: 2, n: 4, a: 2, b: 3, c: 4, want: 6},
		{id: 3, n: 5, a: 2, b: 11, c: 13, want: 10},
		{id: 4, n: 1000000000, a: 2, b: 217983653, c: 336916467, want: 1999999984},
	}
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := nthUglyNumber(tt.n, tt.a, tt.b, tt.c)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
