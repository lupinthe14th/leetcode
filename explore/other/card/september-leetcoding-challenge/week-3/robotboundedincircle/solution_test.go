package robotboundedincircle

import (
	"fmt"
	"testing"
)

func TestIsRobotBounded(t *testing.T) {
	t.Parallel()
	type in struct {
		instructions, m string
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{instructions: "GGLLGG"}, want: true},
		{in: in{instructions: "GG"}, want: false},
		{in: in{instructions: "GL"}, want: true},
		{in: in{instructions: "GLGLGGLGL"}, want: false},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := isRobotBounded(tt.in.instructions)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
