package climbing_stairs

// ClimbStairs returns the number of distinct ways to climb n stairs.
// Uses dynamic programming with O(1) space complexity.
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev2 := 1 // Ways to reach step 1
	prev1 := 2 // Ways to reach step 2

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}
