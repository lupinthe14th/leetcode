package validpalindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   string
		want bool
	}{
		{in: "A man, a plan, a canal: Panama", want: true},
		{in: "race a car", want: false},
		{in: "0P", want: false},
		{in: "", want: true},
		{in: "a", want: true},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPalindrome(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
