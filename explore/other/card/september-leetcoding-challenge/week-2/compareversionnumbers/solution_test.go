package compareversionnumbers

import (
	"fmt"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	t.Parallel()
	type in struct {
		version1, version2 string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{version1: "0.1", version2: "1.1"}, want: -1},
		{in: in{version1: "1.0.1", version2: "1"}, want: 1},
		{in: in{version1: "7.5.2.4", version2: "7.5.3"}, want: -1},
		{in: in{version1: "1.01", version2: "1.001"}, want: 0},
		{in: in{version1: "1.0", version2: "1.0.0"}, want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := compareVersion(tt.in.version1, tt.in.version2)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
