package coursescheduleii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {
	t.Parallel()
	type in struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{numCourses: 2, prerequisites: [][]int{{1, 0}}}, want: []int{0, 1}},
		{in: in{numCourses: 4, prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, want: []int{0, 1, 2, 3}},
		{in: in{numCourses: 3, prerequisites: [][]int{{1, 0}, {1, 2}, {0, 1}}}, want: []int{}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findOrder(tt.in.numCourses, tt.in.prerequisites)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
