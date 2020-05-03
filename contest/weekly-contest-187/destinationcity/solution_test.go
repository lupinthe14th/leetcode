package destinationcity

import (
	"fmt"
	"testing"
)

func TestDestCity(t *testing.T) {
	tests := []struct {
		in   [][]string
		want string
	}{
		{in: [][]string{
			{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"},
		}, want: "Sao Paulo"},
		{in: [][]string{
			{"B", "C"}, {"D", "B"}, {"C", "A"},
		}, want: "A"},
		{in: [][]string{
			{"A", "Z"},
		}, want: "Z"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := destCity(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
