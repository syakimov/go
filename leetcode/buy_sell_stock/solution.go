package stock

func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	buy := prices[0]
	sell := -1 << 31
	profit := 0

	for i := 1; i < len(prices); i++ {
		curr := prices[i]

		if curr-buy > profit { // There is more lucrative sell
			sell = curr
			profit = sell - buy
		} else if buy > curr { // There is more lucrative buy
			buy = curr
		}
	}

	return profit
}
