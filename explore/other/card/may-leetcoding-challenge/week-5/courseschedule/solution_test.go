package courseschedule

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {
	type in struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{numCourses: 2, prerequisites: [][]int{{1, 0}}}, want: true},
		{in: in{numCourses: 2, prerequisites: [][]int{{1, 0}, {0, 1}}}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := canFinish(tt.in.numCourses, tt.in.prerequisites)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
