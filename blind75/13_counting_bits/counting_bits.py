from typing import List


def count_bits(n: int) -> List[int]:
    """
    Returns an array where each element contains the number of 1 bits
    for the corresponding index. Uses dynamic programming with bit manipulation.

    Args:
        n: Non-negative integer

    Returns:
        List of counts of 1 bits for each number from 0 to n
    """
    dp = [0] * (n + 1)

    for i in range(1, n + 1):
        # dp[i] = count of 1 bits in i
        # i >> 1 removes the last bit
        # i & 1 gives us the last bit (0 or 1)
        dp[i] = dp[i >> 1] + (i & 1)

    return dp
