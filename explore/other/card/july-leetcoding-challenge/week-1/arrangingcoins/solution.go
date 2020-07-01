package arrangingcoins

import "math"

func arrangeCoins(n int) int {
	return int(math.Sqrt(2*float64(n)+0.25) - 0.5)
}
