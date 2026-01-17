package best_time_to_buy_and_sell_stock

// MaxProfit returns the maximum profit from buying and selling stock once.
// Uses single pass with O(n) time and O(1) space.
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
