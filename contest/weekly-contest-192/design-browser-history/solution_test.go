package designbrowserhistory

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		act, url string
		steps    int
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{act: "visit", url: "google.com"}},
		{in: in{act: "visit", url: "facebook.com"}},
		{in: in{act: "visit", url: "youtube.com"}},
		{in: in{act: "back", steps: 1}, want: "facebook.com"},
		{in: in{act: "back", steps: 1}, want: "google.com"},
		{in: in{act: "forward", steps: 1}, want: "facebook.com"},
		{in: in{act: "visit", url: "linkedin.com"}},
		{in: in{act: "forward", steps: 2}, want: "linkedin.com"},
		{in: in{act: "back", steps: 2}, want: "google.com"},
		{in: in{act: "back", steps: 7}, want: "leetcode.com"},
	}
	b := Constructor("leetcode.com")
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var got string
			switch tt.in.act {
			case "visit":
				b.Visit(tt.in.url)
			case "back":
				got = b.Back(tt.in.steps)
			case "forward":
				got = b.Forward(tt.in.steps)
			}
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
