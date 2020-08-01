package detectcapital

import (
	"fmt"
	"testing"
)

func TestDetectCapitalUse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   string
		want bool
	}{
		{in: "USA", want: true},
		{in: "leetcode", want: true},
		{in: "Google", want: true},
		{in: "FlaG", want: false},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := detectCapitalUse(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}

}
