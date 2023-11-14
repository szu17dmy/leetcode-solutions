package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	minPrice := prices[0]
	if len(prices) == 1 {
		return 0
	}
	maxProf := prices[1] - prices[0]
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if price-minPrice > maxProf {
			maxProf = price - minPrice
		}
	}
	return maxProf
}
