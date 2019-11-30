package plusone

func plusOne(digits []int) []int {

	var carry int
	for i := len(digits) - 1; i >= 0; i-- {
		var calc int
		if i == len(digits)-1 {
			calc = carry + digits[i] + 1
		} else {
			calc = carry + digits[i]
		}
		digits[i] = calc % 10
		carry = calc / 10
		if carry == 0 {
			break
		}
	}
	if carry != 0 {
		carrys := []int{carry}
		temp := make([]int, len(digits)+1)
		temp = append(carrys, digits...)
		return temp
	}

	return digits
}
