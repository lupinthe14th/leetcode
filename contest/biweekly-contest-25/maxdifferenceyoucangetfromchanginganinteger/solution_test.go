package maxdifferenceyoucangetfromchanginganinteger

import (
	"fmt"
	"testing"
)

func TestMaxDiff(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{in: 555, want: 888},
		{in: 9, want: 8},
		{in: 123456, want: 820000},
		{in: 10000, want: 80000},
		{in: 9288, want: 8700},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxDiff(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
