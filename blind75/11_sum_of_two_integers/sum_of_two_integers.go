package sum_of_two_integers

// GetSum returns the sum of two integers without using + or - operators.
// Uses bit manipulation with XOR for sum and AND for carry.
func GetSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}

	return a
}
