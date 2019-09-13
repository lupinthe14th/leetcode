package reverseinteger

import "math"

func reverse(x int) int {
	r := x
	if x < int(math.Pow(2, 31))*-1 || x > int(math.Pow(2, 31)-1) {
		return 0
	}
	if x < 0 {
		r = r * -1
	}
	reversedNumber := 0
	for i := 0; i < 3; i++ {
		reversedNumber = reversedNumber*10 + r%10
		r /= 10
		if r == 0 {
			break
		}
	}
	if x < 0 {
		return reversedNumber * -1
	}
	return reversedNumber
}
