package coin_change

// CoinChange returns the minimum number of coins needed to make up the amount.
// Returns -1 if the amount cannot be made up.
func CoinChange(coins []int, amount int) int {
	// dp[i] represents minimum coins needed for amount i
	dp := make([]int, amount+1)

	// Initialize with amount+1 (represents impossible/infinity)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	// Base case: 0 coins needed for amount 0
	dp[0] = 0

	// Build up solutions for each amount
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	// If dp[amount] is still infinity, amount cannot be made
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
