package counting_bits

// CountBits returns an array where each element contains the number of 1 bits
// for the corresponding index. Uses dynamic programming with bit manipulation.
func CountBits(n int) []int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		// dp[i] = count of 1 bits in i
		// i >> 1 removes the last bit
		// i & 1 gives us the last bit (0 or 1)
		dp[i] = dp[i>>1] + (i & 1)
	}

	return dp
}
