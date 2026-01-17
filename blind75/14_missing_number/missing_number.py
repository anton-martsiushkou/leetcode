from typing import List


def missing_number(nums: List[int]) -> int:
    """
    Finds the missing number in the range [0, n].
    Uses XOR bit manipulation for O(1) space complexity.

    Args:
        nums: List of distinct integers in range [0, n]

    Returns:
        The missing number in the range
    """
    result = len(nums)

    for i, num in enumerate(nums):
        result ^= i ^ num

    return result
