package wordbreakii

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestWordBreak(t *testing.T) {
	t.Parallel()
	type in struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		in   in
		want []string
	}{
		{in: in{s: "catsanddog", wordDict: []string{"cat", "cats", "and", "sand", "dog"}}, want: []string{"cats and dog", "cat sand dog"}},
		{in: in{s: "pineapplepenapple", wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"}}, want: []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}},
		{in: in{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}}, want: []string{}},
	}

	opts := []cmp.Option{
		cmpopts.SortSlices(func(x, y string) bool { return x < y }),
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := wordBreak(tt.in.s, tt.in.wordDict)
			if !cmp.Equal(got, tt.want, opts...) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
