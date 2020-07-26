package allpathsfromsourcetotarget

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAllPathSourceTarget(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want [][]int
	}{
		{in: [][]int{{1, 2}, {3}, {3}, {}}, want: [][]int{{0, 1, 3}, {0, 2, 3}}},
		{in: [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}, want: [][]int{{0, 3, 6, 7}, {0, 3, 4, 7}, {0, 3, 4, 6, 7}, {0, 3, 4, 5, 6, 7}, {0, 1, 4, 7}, {0, 1, 4, 6, 7}, {0, 1, 4, 5, 6, 7}, {0, 1, 6, 7}, {0, 1, 7}, {0, 1, 2, 4, 7}, {0, 1, 2, 4, 6, 7}, {0, 1, 2, 4, 5, 6, 7}, {0, 1, 2, 6, 7}, {0, 1, 2, 3, 6, 7}, {0, 1, 2, 3, 4, 7}, {0, 1, 2, 3, 4, 6, 7}, {0, 1, 2, 3, 4, 5, 6, 7}, {0, 1, 5, 6, 7}}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := allPathsSourceTarget(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
