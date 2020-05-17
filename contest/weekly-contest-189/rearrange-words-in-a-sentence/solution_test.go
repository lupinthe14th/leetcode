package rearrangewordsinasentence

import (
	"fmt"
	"testing"
)

func TestArrangeWords(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in: "Leetcode is cool", want: "Is cool leetcode"},
		{in: "Keep calm and code on", want: "On and keep calm code"},
		{in: "To be or not to be", want: "To be or to be not"},
		{in: "Jtik hrzrvhbmk gbo cfhmiqwonj ojkew ufvledv bomoeqt ops jgi zdhvbpbb zczmepdfpm jry rjazc titttcu", want: "Gbo ops jgi jry jtik ojkew rjazc ufvledv bomoeqt titttcu zdhvbpbb hrzrvhbmk cfhmiqwonj zczmepdfpm"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := arrangeWords(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
