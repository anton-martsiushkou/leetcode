def climb_stairs(n: int) -> int:
    """
    Returns the number of distinct ways to climb n stairs.
    Uses dynamic programming with O(1) space complexity.

    Args:
        n: Number of stairs to climb

    Returns:
        Number of distinct ways to reach the top
    """
    if n <= 2:
        return n

    prev2 = 1  # Ways to reach step 1
    prev1 = 2  # Ways to reach step 2

    for i in range(3, n + 1):
        current = prev1 + prev2
        prev2 = prev1
        prev1 = current

    return prev1
