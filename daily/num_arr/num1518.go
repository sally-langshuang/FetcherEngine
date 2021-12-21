package num_arr

func numWaterBottles(numBottles int, numExchange int) (ans int) {
	for x := numBottles; ; {
		ans += x / numExchange * numExchange
		x = x%numExchange + x/numExchange
		if x/numExchange == 0 {
			ans += x
			break
		}
	}
	return ans
}
