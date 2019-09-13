package reverseinteger

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	r := x
	if x < int(math.Pow(2, 31))*-1 || x > int(math.Pow(2, 31)-1) {
		return 0
	}
	if x < 0 {
		r = r * -1
	}
	reversedNumber := 0
	l := len(strconv.Itoa(r))
	for i := 0; i < l; i++ {
		reversedNumber = reversedNumber*10 + r%10
		r /= 10
	}
	if x < 0 {
		return reversedNumber * -1
	}
	return reversedNumber
}
