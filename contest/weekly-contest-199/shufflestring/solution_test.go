package shufflestring

import (
	"fmt"
	"testing"
)

func TestRestoreString(t *testing.T) {
	t.Parallel()
	type in struct {
		s       string
		indices []int
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{s: "codeleet", indices: []int{4, 5, 6, 7, 0, 2, 1, 3}}, want: "leetcode"},
		{in: in{s: "abc", indices: []int{0, 1, 2}}, want: "abc"},
		{in: in{s: "aiohn", indices: []int{3, 1, 4, 2, 0}}, want: "nihao"},
		{in: in{s: "aaiougrt", indices: []int{4, 0, 2, 6, 7, 3, 1, 5}}, want: "arigatou"},
		{in: in{s: "art", indices: []int{1, 0, 2}}, want: "rat"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := restoreString(tt.in.s, tt.in.indices)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want %v", tt.in, got, tt.want)
			}
		})
	}
}
