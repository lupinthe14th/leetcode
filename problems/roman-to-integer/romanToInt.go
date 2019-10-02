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

	chars := strings.Split(s, "")
	var result int
	cMap := make(map[int]string)
	for i, c := range chars {
		cMap[i] = c
	}

	for i := 0; i < len(s); i++ {
		ss := cMap[i] + cMap[i+1] // ss is subtraction symbol
		sn, ok := roman[ss]
		if ok {
			result = result + sn
			i++
		} else {
			result = result + roman[cMap[i]]
		}
	}
	return result
}
