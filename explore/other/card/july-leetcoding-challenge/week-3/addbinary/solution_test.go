package addbinary

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	t.Parallel()
	type in struct {
		a, b string
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{a: "11", b: "1"}, want: "100"},
		{in: in{a: "1010", b: "1011"}, want: "10101"},
		{in: in{a: "1", b: "111"}, want: "1000"},
		{in: in{a: "100", b: "110010"}, want: "110110"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := addBinary(tt.in.a, tt.in.b)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
