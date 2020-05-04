package ransomnote

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	type in struct {
		ransomNote, magazine string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{ransomNote: "a", magazine: "b"}, want: false},
		{in: in{ransomNote: "aa", magazine: "ab"}, want: false},
		{in: in{ransomNote: "aa", magazine: "aab"}, want: true},
		{in: in{ransomNote: "", magazine: ""}, want: true},
		{in: in{ransomNote: "bg", magazine: "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"}, want: true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := canConstruct(tt.in.ransomNote, tt.in.magazine)
			if got != tt.want {
				t.Fatalf("in: %#v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
