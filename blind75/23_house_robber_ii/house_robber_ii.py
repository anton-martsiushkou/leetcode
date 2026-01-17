from typing import List


def rob(nums: List[int]) -> int:
    """
    Returns the maximum amount of money that can be robbed from circular houses.
    Uses dynamic programming with two passes: O(n) time and O(1) space complexity.

    Args:
        nums: List of integers representing money in each house

    Returns:
        Maximum amount of money that can be robbed
    """
    n = len(nums)
    if n == 0:
        return 0
    if n == 1:
        return nums[0]
    if n == 2:
        return max(nums[0], nums[1])

    # Two scenarios: rob houses [0...n-2] or [1...n-1]
    return max(rob_range(nums, 0, n - 2), rob_range(nums, 1, n - 1))


def rob_range(nums: List[int], start: int, end: int) -> int:
    """
    Applies the House Robber I algorithm to a range of houses.

    Args:
        nums: List of integers representing money in each house
        start: Starting index
        end: Ending index (inclusive)

    Returns:
        Maximum amount of money that can be robbed in the range
    """
    prev2 = 0
    prev1 = 0

    for i in range(start, end + 1):
        current = max(prev1, prev2 + nums[i])
        prev2 = prev1
        prev1 = current

    return prev1
