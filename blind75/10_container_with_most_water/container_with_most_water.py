from typing import List


def max_area(height: List[int]) -> int:
    """
    Finds the maximum area of water that can be contained.
    Uses two-pointer approach with O(n) time complexity.

    Args:
        height: List of heights representing vertical lines

    Returns:
        Maximum area of water that can be contained
    """
    left, right = 0, len(height) - 1
    max_area_val = 0

    while left < right:
        # Calculate area with current two lines
        width = right - left
        h = min(height[left], height[right])
        area = width * h

        # Update maximum area
        max_area_val = max(max_area_val, area)

        # Move the pointer with shorter height
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1

    return max_area_val
