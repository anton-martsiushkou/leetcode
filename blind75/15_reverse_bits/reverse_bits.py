def reverse_bits(n: int) -> int:
    """
    Reverses the bits of a 32-bit unsigned integer.
    Uses bit-by-bit reversal for clarity and O(1) space.

    Args:
        n: A 32-bit unsigned integer

    Returns:
        The integer with reversed bits
    """
    result = 0

    for i in range(32):
        result <<= 1       # Shift result left to make room
        result |= (n & 1)  # Add the rightmost bit of n
        n >>= 1            # Move to next bit in n

    return result
