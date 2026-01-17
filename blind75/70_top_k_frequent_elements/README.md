# Top K Frequent Elements

## Problem Description

Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in any order.

## Examples

**Example 1:**
```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

**Example 2:**
```
Input: nums = [1], k = 1
Output: [1]
```

## Constraints

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `k` is in the range `[1, the number of unique elements in the array]`
- It is guaranteed that the answer is unique.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Bucket Sort | Hash Map + Array | O(n) | O(n) |
| Min Heap | Hash Map + Min Heap | O(n log k) | O(n + k) |

### Approach 1: Bucket Sort (Optimal)

The optimal solution uses bucket sort to achieve O(n) time complexity. We first count the frequency of each element, then create buckets where the index represents the frequency. Finally, we collect elements from the highest frequency buckets.

**Key Insight:** The maximum frequency of any element cannot exceed the array length, so we can use an array of size n+1 where each index represents a frequency.

### Algorithm Steps

1. Count the frequency of each element using a hash map
2. Create an array of buckets (size n+1) where each bucket stores elements with that frequency
3. Iterate from the highest frequency bucket to the lowest
4. Collect elements until we have k elements

### Example Walkthrough

For `nums = [1,1,1,2,2,3]` and `k = 2`:

1. **Count frequencies**:
   - `{1: 3, 2: 2, 3: 1}`

2. **Create buckets**:
   - `buckets[1] = [3]`
   - `buckets[2] = [2]`
   - `buckets[3] = [1]`

3. **Collect from highest frequency**:
   - From bucket[3]: add 1
   - From bucket[2]: add 2
   - Result: `[1, 2]`

### Approach 2: Min Heap

Use a min heap of size k to keep track of the k most frequent elements. This approach is useful when k is small compared to n.

### Why Bucket Sort is Optimal

- **Time Complexity O(n)**: We traverse the array once for counting, once for bucketing, and once for collecting results
- **Space Complexity O(n)**: We need space for the frequency map and buckets
- This is optimal because we must examine every element at least once to count frequencies

### Alternative Approaches

1. **Sort by Frequency**: O(n log n) time - Slower than bucket sort
2. **Max Heap**: O(n log n) time - Build heap with all elements, then extract k times
