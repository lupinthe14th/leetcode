package bullsandcows

import (
	"fmt"
	"testing"
)

func TestGetHint(t *testing.T) {
	t.Parallel()
	type in struct {
		secret, guess string
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{secret: "1807", guess: "7810"}, want: "1A3B"},
		{in: in{secret: "1123", guess: "0111"}, want: "1A1B"},
		{in: in{secret: "11", guess: "10"}, want: "1A0B"},
		{in: in{secret: "11", guess: "01"}, want: "1A0B"},
		{in: in{secret: "11", guess: "11"}, want: "2A0B"},
		{in: in{secret: "1122", guess: "2211"}, want: "0A4B"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := getHint(tt.in.secret, tt.in.guess)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
