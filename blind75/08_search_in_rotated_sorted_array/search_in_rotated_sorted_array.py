from typing import List


def search(nums: List[int], target: int) -> int:
    """
    Finds the index of target in a rotated sorted array.
    Uses modified binary search with O(log n) time complexity.

    Args:
        nums: Rotated sorted array with unique elements
        target: Target value to search for

    Returns:
        Index of target if found, -1 otherwise
    """
    left, right = 0, len(nums) - 1

    while left <= right:
        mid = left + (right - left) // 2

        if nums[mid] == target:
            return mid

        # Determine which half is sorted
        if nums[left] <= nums[mid]:
            # Left half is sorted
            if nums[left] <= target < nums[mid]:
                # Target is in the sorted left half
                right = mid - 1
            else:
                # Target is in the right half
                left = mid + 1
        else:
            # Right half is sorted
            if nums[mid] < target <= nums[right]:
                # Target is in the sorted right half
                left = mid + 1
            else:
                # Target is in the left half
                right = mid - 1

    return -1
