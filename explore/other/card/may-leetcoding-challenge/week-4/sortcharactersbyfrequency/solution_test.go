package sortcharactersbyfrequency

import (
	"fmt"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in: "tree", want: "eert"},
		{in: "cccaaa", want: "aaaccc"},
		{in: "Aabb", want: "bbAa"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := frequencySort(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
