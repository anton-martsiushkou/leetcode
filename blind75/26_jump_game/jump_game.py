from typing import List


def can_jump(nums: List[int]) -> bool:
    """
    Determines if it's possible to reach the last index.
    Uses a greedy algorithm with O(n) time and O(1) space complexity.

    Args:
        nums: List of integers representing maximum jump length at each position

    Returns:
        True if we can reach the last index, False otherwise
    """
    max_reach = 0
    n = len(nums)

    for i in range(n):
        # If current position is beyond max reachable, can't proceed
        if i > max_reach:
            return False

        # Update the farthest position we can reach
        max_reach = max(max_reach, i + nums[i])

        # If we can reach the last index, return True
        if max_reach >= n - 1:
            return True

    return True
