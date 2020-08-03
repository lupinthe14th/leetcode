package designhashset

import (
	"fmt"
	"testing"
)

func TestSumRange(t *testing.T) {
	type in struct {
		act string
		key int
	}

	tests := []struct {
		in   in
		want bool
	}{
		{in: in{act: "add", key: 1}},
		{in: in{act: "add", key: 2}},
		{in: in{act: "cont", key: 1}, want: true},
		{in: in{act: "cont", key: 3}, want: false},
		{in: in{act: "add", key: 2}},
		{in: in{act: "cont", key: 2}, want: true},
		{in: in{act: "rem", key: 2}},
		{in: in{act: "cont", key: 2}, want: false},
	}

	c := Constructor()
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			switch tt.in.act {
			case "add":
				c.Add(tt.in.key)
			case "cont":
				got := c.Contains(tt.in.key)
				if got != tt.want {
					t.Fatalf("in %v got: %v want: %v", tt.in, got, tt.want)
				}
			case "rem":
				c.Remove(tt.in.key)
			}
		})
	}
}
