package reversewordsinastring

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want string
	}{
		{in: "the sky is blue", want: "blue is sky the"},
		{in: "  hello world!  ", want: "world! hello"},
		{in: "a good   example", want: "example good a"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := reverseWords(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
