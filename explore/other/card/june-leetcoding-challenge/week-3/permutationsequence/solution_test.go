package permutationsequence

import (
	"fmt"
	"testing"
)

func TestGetPermutation(t *testing.T) {
	type in struct {
		n, k int
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{n: 3, k: 3}, want: "213"},
		{in: in{n: 4, k: 9}, want: "2314"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := getPermutation(tt.in.n, tt.in.k)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
