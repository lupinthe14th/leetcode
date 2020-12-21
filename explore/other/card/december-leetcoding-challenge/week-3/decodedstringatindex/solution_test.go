package decodedstringatindex

import (
	"fmt"
	"testing"
)

func TestDecodeAtIndex(t *testing.T) {
	t.Parallel()
	type in struct {
		S string
		K int
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{S: "leet2code3", K: 10}, want: "o"},
		{in: in{S: "ha22", K: 5}, want: "h"},
		{in: in{S: "a2345678999999999999999", K: 1}, want: "a"},
		{in: in{S: "y959q969u3hb22odq595", K: 222280369}, want: "y"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := decodeAtIndex(tt.in.S, tt.in.K)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
