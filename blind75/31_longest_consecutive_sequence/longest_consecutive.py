from typing import List


def longest_consecutive(nums: List[int]) -> int:
    """
    Finds the longest consecutive sequence using hash set.
    Time: O(n), Space: O(n)

    Args:
        nums: List of integers

    Returns:
        Length of longest consecutive sequence
    """
    if not nums:
        return 0

    # Build hash set
    num_set = set(nums)
    max_length = 0

    # Check each number
    for num in num_set:
        # Only start counting if this is the beginning of a sequence
        if num - 1 not in num_set:
            current_num = num
            current_length = 1

            # Count consecutive numbers
            while current_num + 1 in num_set:
                current_num += 1
                current_length += 1

            max_length = max(max_length, current_length)

    return max_length


def longest_consecutive_sorting(nums: List[int]) -> int:
    """
    Finds the longest consecutive sequence using sorting.
    Time: O(n log n), Space: O(1) or O(n) depending on sort

    Args:
        nums: List of integers

    Returns:
        Length of longest consecutive sequence
    """
    if not nums:
        return 0

    nums.sort()
    max_length = 1
    current_length = 1

    for i in range(1, len(nums)):
        if nums[i] == nums[i - 1]:
            # Skip duplicates
            continue
        elif nums[i] == nums[i - 1] + 1:
            # Consecutive
            current_length += 1
            max_length = max(max_length, current_length)
        else:
            # Break in sequence
            current_length = 1

    return max_length
