package reconstructitinerary

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	tests := []struct {
		in   [][]string
		want []string
	}{
		{in: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}, want: []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		{in: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}, want: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
		{in: [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}, want: []string{"JFK", "NRT", "JFK", "KUL"}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findItinerary(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
