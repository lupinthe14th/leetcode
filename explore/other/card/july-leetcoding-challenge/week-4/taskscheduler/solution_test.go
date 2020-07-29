package taskscheduler

import (
	"fmt"
	"testing"
)

func TestLeastInterval(t *testing.T) {
	t.Parallel()
	type in struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 2}, want: 8},
		{in: in{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 3}, want: 10},
		{in: in{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 0}, want: 6},
		{in: in{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 50}, want: 104},
		{in: in{tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, n: 2}, want: 16},
		{in: in{tasks: []byte{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D', 'E', 'E', 'F', 'F', 'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J', 'K', 'K', 'L', 'L', 'M', 'M', 'N', 'N', 'O', 'O', 'P', 'P', 'Q', 'Q', 'R', 'R', 'S', 'S', 'T', 'T', 'U', 'U', 'V', 'V', 'W', 'W', 'X', 'X', 'Y', 'Y', 'Z', 'Z'}, n: 2}, want: 52},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			// t.Parallel()
			got := leastInterval(tt.in.tasks, tt.in.n)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
