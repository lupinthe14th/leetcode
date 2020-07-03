package averagesalaryexcludingtheminimumandmaximumsalary

import (
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		in   []int
		want float64
	}{
		{in: []int{4000, 3000, 1000, 2000}, want: 2500.00000},
		{in: []int{1000, 2000, 3000}, want: 2000.00000},
		{in: []int{6000, 5000, 4000, 3000, 2000, 1000}, want: 3500.00000},
		{in: []int{8000, 9000, 2000, 3000, 6000, 1000}, want: 4750.00000},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := average(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
