# Non-overlapping Intervals

## Problem Description

Given an array of intervals `intervals` where `intervals[i] = [starti, endi]`, return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

## Examples

**Example 1:**
```
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.
```

**Example 2:**
```
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.
```

**Example 3:**
```
Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
```

**Example 4:**
```
Input: intervals = [[1,100],[11,22],[1,11],[2,12]]
Output: 2
```

## Constraints

- `1 <= intervals.length <= 10^5`
- `intervals[i].length == 2`
- `-5 * 10^4 <= starti < endi <= 5 * 10^4`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Greedy (Sort by End Time) | Array | O(n log n) | O(1) |

### Approach: Greedy Algorithm

This problem is equivalent to the classic "Activity Selection Problem". The key insight is that we want to keep as many intervals as possible (which minimizes removals). To maximize the number of non-overlapping intervals, we should greedily choose intervals that end earliest, as they leave the most room for future intervals.

**Key Insight:** Sort intervals by end time and greedily select intervals that don't overlap. The number of intervals we remove is the total count minus the number we can keep.

### Algorithm Steps

1. **Sort intervals by end time** (second element of each interval)
2. Initialize `count` = 1 (we can always keep the first interval after sorting)
3. Track `end` = end time of the last kept interval (start with first interval's end)
4. **Iterate through remaining intervals:**
   - If current interval's start ≥ previous end (no overlap):
     - Keep this interval: increment `count`
     - Update `end` to current interval's end
   - Else (overlap):
     - Skip this interval (implicitly remove it)
5. Return `total intervals - count` (number of removals)

### Example Walkthrough

For `intervals = [[1,2],[2,3],[3,4],[1,3]]`:

1. **Sort by end time:**
   - `[[1,2],[2,3],[1,3],[3,4]]`

2. **Initialize:**
   - `count = 1` (keep first interval [1,2])
   - `end = 2`

3. **Process [2,3]:**
   - start = 2 ≥ end = 2 (no overlap)
   - Keep it: `count = 2`, `end = 3`

4. **Process [1,3]:**
   - start = 1 < end = 3 (overlap)
   - Remove it (don't update count or end)

5. **Process [3,4]:**
   - start = 3 ≥ end = 3 (no overlap)
   - Keep it: `count = 3`, `end = 4`

6. **Result:** Total (4) - Kept (3) = **1 removal**

### Why This Greedy Approach Works

The greedy choice of always selecting the interval that ends earliest is optimal because:

1. **Leaves Maximum Space:** An interval that ends earlier leaves more room for future intervals
2. **Optimal Substructure:** After making the greedy choice, the remaining problem has the same structure
3. **No Better Alternative:** Choosing an interval that ends later can never be better, as it reduces the space for future intervals

### Why This is Optimal

- **Time Complexity O(n log n)**: Dominated by the sorting step. The iteration is O(n)
- **Space Complexity O(1)**: Only uses a constant amount of extra space (ignoring the space used by sorting, which is typically O(log n) for the call stack)
- This is optimal because we must examine every interval at least once, and comparison-based sorting cannot be faster than O(n log n)

### Edge Cases

1. **No overlapping intervals**: Return 0 (no removals needed)
2. **All intervals overlap**: Return n-1 (keep only one)
3. **Identical intervals**: Return n-1 (keep only one)
4. **Single interval**: Return 0 (no removals needed)
5. **Nested intervals**: The greedy algorithm handles this correctly by choosing the one that ends earliest
6. **Touching intervals** (end of one equals start of next): Not considered overlapping

### Alternative Approach: Dynamic Programming

You could also solve this with DP (similar to Longest Increasing Subsequence), but it would be O(n²) time, which is less efficient than the greedy approach.
