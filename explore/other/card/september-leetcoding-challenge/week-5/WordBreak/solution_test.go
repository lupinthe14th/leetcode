package wordbreak

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	t.Parallel()
	type in struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{s: "leetcode", wordDict: []string{"leet", "code"}}, want: true},
		{in: in{s: "applepenapple", wordDict: []string{"apple", "pen"}}, want: true},
		{in: in{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}}, want: false},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := wordBreak(tt.in.s, tt.in.wordDict)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
