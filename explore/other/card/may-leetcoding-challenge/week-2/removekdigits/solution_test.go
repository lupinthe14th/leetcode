package removekdigits

import (
	"fmt"
	"testing"
)

func TestRemoveKdigits(t *testing.T) {
	type in struct {
		num string
		k   int
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{num: "1432219", k: 3}, want: "1219"},
		{in: in{num: "10200", k: 1}, want: "200"},
		{in: in{num: "10", k: 1}, want: "0"},
		{in: in{num: "10", k: 2}, want: "0"},
		{in: in{num: "112", k: 1}, want: "11"},
		{in: in{num: "5337", k: 2}, want: "33"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := removeKdigits(tt.in.num, tt.in.k)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)

			}
		})
	}
}
