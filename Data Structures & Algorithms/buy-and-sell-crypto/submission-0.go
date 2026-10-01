func maxProfit(prices []int) int {
	max := 0
	if len(prices) == 1 {
		return max
	}

	buyDay := 0
	for sellDay := 1; sellDay < len(prices); sellDay++ {
		if prices[sellDay] < prices[buyDay] {
			buyDay = sellDay
		}

		profit := prices[sellDay] - prices[buyDay]
		max = int(math.Max(float64(profit), float64(max)))
	}

	return max
}
