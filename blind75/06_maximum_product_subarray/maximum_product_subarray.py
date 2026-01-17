from typing import List


def max_product(nums: List[int]) -> int:
    """
    Finds the contiguous subarray with the largest product.
    Uses modified Kadane's algorithm tracking both max and min products.

    Args:
        nums: List of integers

    Returns:
        The maximum product of a contiguous subarray
    """
    if not nums:
        return 0

    max_so_far = nums[0]
    max_ending_here = nums[0]
    min_ending_here = nums[0]

    for i in range(1, len(nums)):
        num = nums[i]

        # Calculate candidates: current number, max*num, min*num
        # We need temp variable because max_ending_here is used in min_ending_here calculation
        temp_max = max(num, num * max_ending_here, num * min_ending_here)
        min_ending_here = min(num, num * max_ending_here, num * min_ending_here)
        max_ending_here = temp_max

        max_so_far = max(max_so_far, max_ending_here)

    return max_so_far
