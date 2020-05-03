package lrucache

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	type in struct {
		action     string
		key, value int
	}
	cache := Constructor(2)
	tests := []struct {
		in   in
		want int
	}{
		{in: in{action: "p", key: 2, value: 1}},
		{in: in{action: "p", key: 1, value: 1}},
		{in: in{action: "p", key: 2, value: 3}},
		{in: in{action: "p", key: 4, value: 1}},
		{in: in{action: "g", key: 1}, want: -1},
		{in: in{action: "g", key: 2}, want: 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			switch tt.in.action {
			case "p":
				cache.Put(tt.in.key, tt.in.value)
			case "g":
				got := cache.Get(tt.in.key)
				if got != tt.want {
					t.Fatalf("got: %v, want: %v", got, tt.want)
				}
			}
		})
	}

}
