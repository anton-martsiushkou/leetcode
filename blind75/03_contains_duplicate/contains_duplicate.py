from typing import List


def contains_duplicate(nums: List[int]) -> bool:
    """
    Returns true if any value appears at least twice in the array.
    Uses hash set for O(n) time complexity.

    Args:
        nums: List of integers

    Returns:
        True if any value appears at least twice, False otherwise
    """
    seen = set()

    for num in nums:
        if num in seen:
            return True
        seen.add(num)

    return False
