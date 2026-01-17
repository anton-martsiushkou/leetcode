from typing import List


def top_k_frequent(nums: List[int], k: int) -> List[int]:
    """
    Returns the k most frequent elements using bucket sort.
    Time: O(n), Space: O(n)

    Args:
        nums: List of integers
        k: Number of most frequent elements to return

    Returns:
        List of k most frequent elements
    """
    # Count frequencies
    freq = {}
    for num in nums:
        freq[num] = freq.get(num, 0) + 1

    # Create buckets where index is frequency
    buckets = [[] for _ in range(len(nums) + 1)]
    for num, count in freq.items():
        buckets[count].append(num)

    # Collect k most frequent elements from highest frequency buckets
    result = []
    for i in range(len(buckets) - 1, -1, -1):
        if buckets[i]:
            result.extend(buckets[i])
            if len(result) >= k:
                break

    return result[:k]
