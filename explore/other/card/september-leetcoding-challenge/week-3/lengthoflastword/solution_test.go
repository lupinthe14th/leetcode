package lengthoflastword

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	t.Parallel()
	type in struct {
		s string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{s: "Hello World"}, want: 5},
		{in: in{s: "a "}, want: 1},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := lengthOfLastWord(tt.in.s)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
