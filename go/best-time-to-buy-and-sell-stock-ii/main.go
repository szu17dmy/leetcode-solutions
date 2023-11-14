package best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) (p int) {
	for i := 1; i < len(prices); i++ {
		delta := prices[i] - prices[i-1]
		if delta > 0 {
			p += delta
		}
	}
	return p
}
