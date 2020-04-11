package backspacestringcompare

import (
	"fmt"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	type in struct {
		s, t string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{s: "ab#c", t: "ad#c"}, want: true},
		{in: in{s: "ab##", t: "c#d#"}, want: true},
		{in: in{s: "a##c", t: "#a#c"}, want: true},
		{in: in{s: "a#c", t: "b"}, want: false},
		{in: in{s: "a#c", t: ""}, want: false},
		{in: in{s: "xywrrmp", t: "xywrrmu#p"}, want: true},
		{in: in{s: "bbbextm", t: "bbb#extm"}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := backspaceCompare(tt.in.s, tt.in.t)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
