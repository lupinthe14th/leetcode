package waterbottles

func numWaterBottles(numBottles int, numExchange int) int {
	out, emp := 0, 0
	for {
		// Drink
		out += numBottles
		emp += numBottles
		// exchange
		numBottles = emp / numExchange
		emp %= numExchange
		// break
		if numBottles == 0 && emp < numExchange {
			break
		}
	}
	return out
}
