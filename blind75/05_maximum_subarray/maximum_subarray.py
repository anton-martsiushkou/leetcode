from typing import List


def max_sub_array(nums: List[int]) -> int:
    """
    Finds the contiguous subarray with the largest sum.
    Uses Kadane's Algorithm with O(n) time and O(1) space.

    Args:
        nums: List of integers

    Returns:
        The maximum sum of a contiguous subarray
    """
    max_sum = nums[0]
    current_sum = nums[0]

    for i in range(1, len(nums)):
        # Either extend existing subarray or start new one
        current_sum = max(nums[i], current_sum + nums[i])
        # Update global maximum
        max_sum = max(max_sum, current_sum)

    return max_sum
