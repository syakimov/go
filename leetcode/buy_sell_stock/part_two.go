package stock

func maxProfitSimple(prices []int) int {
	p := 0

	// while the price is going up we are winning
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			p += prices[i] - prices[i-1]
		}
	}

	return p
}

// While the price is up we cannot buy
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	profit := 0

	valley := prices[0]
	peak := prices[0]

	for i := 0; i < len(prices)-1; {
		// find a valley    current price is bigger than next price
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]

		// find a peak          current price is lower or equal than next
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}

		peak = prices[i]

		profit += peak - valley
	}

	return profit
}
