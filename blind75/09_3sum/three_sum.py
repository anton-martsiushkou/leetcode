from typing import List


def three_sum(nums: List[int]) -> List[List[int]]:
    """
    Finds all unique triplets that sum to zero.
    Uses sort + two pointers approach with O(n²) time complexity.

    Args:
        nums: List of integers

    Returns:
        List of triplets that sum to zero
    """
    result = []
    n = len(nums)

    # Sort the array to enable two-pointer approach and duplicate handling
    nums.sort()

    for i in range(n - 2):
        # Skip duplicates for the first element
        if i > 0 and nums[i] == nums[i - 1]:
            continue

        # Two-pointer approach for the remaining elements
        left, right = i + 1, n - 1
        target = -nums[i]

        while left < right:
            current_sum = nums[left] + nums[right]

            if current_sum == target:
                result.append([nums[i], nums[left], nums[right]])

                # Skip duplicates for left pointer
                while left < right and nums[left] == nums[left + 1]:
                    left += 1
                # Skip duplicates for right pointer
                while left < right and nums[right] == nums[right - 1]:
                    right -= 1

                left += 1
                right -= 1
            elif current_sum < target:
                left += 1
            else:
                right -= 1

    return result
