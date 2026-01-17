package number_of_1_bits

// HammingWeight returns the number of '1' bits in the binary representation.
// Uses Brian Kernighan's algorithm for optimal performance.
func HammingWeight(n int) int {
	count := 0

	for n != 0 {
		n = n & (n - 1) // Clear the rightmost set bit
		count++
	}

	return count
}
