def hamming_weight(n: int) -> int:
    """
    Returns the number of '1' bits in the binary representation.
    Uses Brian Kernighan's algorithm for optimal performance.

    Args:
        n: A positive integer

    Returns:
        The count of set bits (Hamming weight)
    """
    count = 0

    while n != 0:
        n = n & (n - 1)  # Clear the rightmost set bit
        count += 1

    return count
