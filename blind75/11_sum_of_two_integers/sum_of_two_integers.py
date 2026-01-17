def get_sum(a: int, b: int) -> int:
    """
    Returns the sum of two integers without using + or - operators.
    Uses bit manipulation with XOR for sum and AND for carry.

    Args:
        a: First integer
        b: Second integer

    Returns:
        Sum of a and b
    """
    # Mask to handle 32-bit integers (Python has arbitrary precision)
    mask = 0xFFFFFFFF

    while b != 0:
        # Calculate sum without carry using XOR
        sum_without_carry = (a ^ b) & mask
        # Calculate carry using AND and left shift
        carry = ((a & b) << 1) & mask

        a = sum_without_carry
        b = carry

    # Handle negative numbers (convert from 32-bit two's complement)
    if a > 0x7FFFFFFF:
        return ~(a ^ mask)

    return a
