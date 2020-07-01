package pathcrossing

import (
	"fmt"
	"testing"
)

func TestIsPathCrossing(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{in: "NES", want: false},
		{in: "NESWW", want: true},
		{in: "NNSWWEWSSESSWENNW", want: true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPathCrossing(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
