from typing import List


def unique_paths(m: int, n: int) -> int:
    """
    Returns the number of unique paths from top-left to bottom-right.
    Uses dynamic programming with O(m*n) time and O(n) space complexity.

    Args:
        m: Number of rows in the grid
        n: Number of columns in the grid

    Returns:
        Number of unique paths from top-left to bottom-right
    """
    # Create a 1D array to store the number of paths to each cell in current row
    dp = [1] * n

    # Process each row starting from row 1
    for i in range(1, m):
        for j in range(1, n):
            # dp[j] currently holds paths from above (previous row)
            # dp[j-1] holds paths from left (current row)
            dp[j] = dp[j] + dp[j - 1]

    return dp[n - 1]
