from typing import List


def find_min(nums: List[int]) -> int:
    """
    Finds the minimum element in a rotated sorted array.
    Uses modified binary search with O(log n) time complexity.

    Args:
        nums: Rotated sorted array with unique elements

    Returns:
        The minimum element in the array
    """
    left, right = 0, len(nums) - 1

    # Binary search for the minimum element
    while left < right:
        mid = left + (right - left) // 2

        # If mid element is greater than right element,
        # minimum is in the right half
        if nums[mid] > nums[right]:
            left = mid + 1
        else:
            # Otherwise, minimum is in the left half (including mid)
            right = mid

    # Left and right converge to the minimum element
    return nums[left]
