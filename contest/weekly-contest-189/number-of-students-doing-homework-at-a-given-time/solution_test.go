package numberofstudentsdoinghomeworkatagiventime

import (
	"fmt"
	"testing"
)

func TestBusyStudent(t *testing.T) {
	type in struct {
		startTime, endTime []int
		queryTime          int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{startTime: []int{1, 2, 3}, endTime: []int{3, 2, 7}, queryTime: 4}, want: 1},
		{in: in{startTime: []int{4}, endTime: []int{4}, queryTime: 4}, want: 1},
		{in: in{startTime: []int{4}, endTime: []int{4}, queryTime: 5}, want: 0},
		{in: in{startTime: []int{1, 1, 1, 1}, endTime: []int{1, 3, 2, 4}, queryTime: 7}, want: 0},
		{in: in{startTime: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, endTime: []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, queryTime: 5}, want: 5},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := busyStudent(tt.in.startTime, tt.in.endTime, tt.in.queryTime)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
