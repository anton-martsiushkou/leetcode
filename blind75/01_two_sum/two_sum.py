from typing import List


def two_sum(nums: List[int], target: int) -> List[int]:
    """
    Returns indices of two numbers that add up to target.
    Uses a hash map for O(n) time complexity.

    Args:
        nums: List of integers
        target: Target sum

    Returns:
        List containing indices of the two numbers that add up to target
    """
    seen = {}

    for i, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            return [seen[complement], i]
        seen[num] = i

    return []
