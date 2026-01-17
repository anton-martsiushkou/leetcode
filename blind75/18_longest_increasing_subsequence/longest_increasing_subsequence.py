from typing import List
import bisect


def length_of_lis(nums: List[int]) -> int:
    """
    Returns the length of the longest increasing subsequence.
    Uses dynamic programming with binary search for O(n log n) time complexity.

    Args:
        nums: List of integers

    Returns:
        Length of the longest strictly increasing subsequence
    """
    if not nums:
        return 0

    # tails[i] is the smallest tail of all increasing subsequences of length i+1
    tails = []

    for num in nums:
        # Binary search for the position to insert/replace
        pos = bisect.bisect_left(tails, num)

        # If num is larger than all elements, append it
        if pos == len(tails):
            tails.append(num)
        else:
            # Replace the element at position pos
            tails[pos] = num

    return len(tails)
