package checkifall1sareatleastlengthkplacesaway

import (
	"fmt"
	"testing"
)

func TestKLengthApart(t *testing.T) {
	type in struct {
		nums []int
		k    int
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{nums: []int{1, 0, 0, 0, 1, 0, 0, 1}, k: 2}, want: true},
		{in: in{nums: []int{1, 0, 0, 1, 0, 1}, k: 2}, want: false},
		{in: in{nums: []int{1, 1, 1, 1, 1}, k: 0}, want: true},
		{in: in{nums: []int{0, 1, 0, 1}, k: 1}, want: true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := kLengthApart(tt.in.nums, tt.in.k)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
