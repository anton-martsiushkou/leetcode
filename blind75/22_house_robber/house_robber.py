from typing import List


def rob(nums: List[int]) -> int:
    """
    Returns the maximum amount of money that can be robbed.
    Uses dynamic programming with O(n) time and O(1) space complexity.

    Args:
        nums: List of integers representing money in each house

    Returns:
        Maximum amount of money that can be robbed
    """
    if not nums:
        return 0
    if len(nums) == 1:
        return nums[0]
    if len(nums) == 2:
        return max(nums[0], nums[1])

    # prev2 = max money robbed up to house i-2
    # prev1 = max money robbed up to house i-1
    prev2 = nums[0]
    prev1 = max(nums[0], nums[1])

    for i in range(2, len(nums)):
        current = max(prev1, prev2 + nums[i])
        prev2 = prev1
        prev1 = current

    return prev1
