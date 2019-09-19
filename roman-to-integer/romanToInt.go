package romantointeger

import (
	"strings"
)

func romanToInt(s string) int {
	var roman = map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	var result int
	for len(s) > 0 {
		for k, v := range roman {
			if strings.HasSuffix(s, k) {
				result = result + v
				s = strings.TrimSuffix(s, k)
			}
		}
	}
	return result
}
