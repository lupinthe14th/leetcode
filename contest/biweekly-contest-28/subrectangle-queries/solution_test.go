package subrectanglequeries

import (
	"fmt"
	"testing"
)

func TestSubrectangleQueries(t *testing.T) {
	type in struct {
		act                                        string
		rectangle                                  [][]int
		row, col, row1, col1, row2, col2, newValue int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{act: "get", row: 0, col: 2}, want: 1},
		{in: in{act: "update", row1: 0, col1: 0, row2: 3, col2: 2, newValue: 5}},
		{in: in{act: "get", row: 0, col: 2}, want: 5},
		{in: in{act: "get", row: 3, col: 1}, want: 5},
		{in: in{act: "update", row1: 3, col1: 0, row2: 3, col2: 2, newValue: 10}},
		{in: in{act: "get", row: 3, col: 1}, want: 10},
		{in: in{act: "get", row: 0, col: 2}, want: 5},
	}
	s := Constructor([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var got int
			switch tt.in.act {
			case "get":
				got = s.GetValue(tt.in.row, tt.in.col)
			case "update":
				s.UpdateSubrectangle(tt.in.row1, tt.in.col1, tt.in.row2, tt.in.col2, tt.in.newValue)
			}
			if got != tt.want {
				t.Fatalf("in: %v got :%v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
