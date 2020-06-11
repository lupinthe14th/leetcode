package sortcolors

func sortColors(nums []int) {
	red, white, blue := 0, 0, 0
	for _, n := range nums {
		switch n {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}

	for i := range nums {
		switch {
		case red > 0:
			nums[i] = 0
			red--
		case white > 0:
			nums[i] = 1
			white--
		case blue > 0:
			nums[i] = 2
			blue--
		}
	}
}
