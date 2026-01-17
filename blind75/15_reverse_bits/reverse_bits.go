package reverse_bits

// ReverseBits reverses the bits of a 32-bit unsigned integer.
// Uses bit-by-bit reversal for clarity and O(1) space.
func ReverseBits(n uint32) uint32 {
	var result uint32

	for i := 0; i < 32; i++ {
		result <<= 1        // Shift result left to make room
		result |= (n & 1)   // Add the rightmost bit of n
		n >>= 1             // Move to next bit in n
	}

	return result
}
