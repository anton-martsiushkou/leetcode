from typing import List


def product_except_self(nums: List[int]) -> List[int]:
    """
    Returns an array where each element is the product of all elements except itself.
    Uses prefix-suffix product approach with O(n) time and O(1) extra space.

    Args:
        nums: List of integers

    Returns:
        List where each element is the product of all other elements
    """
    n = len(nums)
    result = [1] * n

    # First pass: fill result with prefix products
    prefix = 1
    for i in range(n):
        result[i] = prefix
        prefix *= nums[i]

    # Second pass: multiply by suffix products
    suffix = 1
    for i in range(n - 1, -1, -1):
        result[i] *= suffix
        suffix *= nums[i]

    return result
