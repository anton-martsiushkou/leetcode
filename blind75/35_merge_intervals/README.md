# Merge Intervals

## Problem Description

Given an array of `intervals` where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

## Examples

**Example 1:**
```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
```

**Example 2:**
```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```

**Example 3:**
```
Input: intervals = [[1,4],[0,4]]
Output: [[0,4]]
```

**Example 4:**
```
Input: intervals = [[1,4],[2,3]]
Output: [[1,4]]
```

## Constraints

- `1 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 10^4`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Sort and Merge | Array | O(n log n) | O(n) |

### Approach: Sort and Merge

The solution first sorts the intervals by their start time, then iterates through them once to merge overlapping intervals.

**Key Insight:** Once intervals are sorted by start time, we only need to compare each interval with the last merged interval. If they overlap, we extend the end time of the last merged interval. If they don't overlap, we add the current interval as a new merged interval.

### Algorithm Steps

1. **Sort intervals** by start time (first element of each interval)
2. Initialize result array with the first interval
3. **Iterate through remaining intervals:**
   - Get the last interval in result
   - If current interval overlaps with last interval (current.start ≤ last.end):
     - Merge by updating last.end = max(last.end, current.end)
   - Else:
     - Add current interval to result (no overlap)
4. Return the result

### Example Walkthrough

For `intervals = [[1,3],[2,6],[8,10],[15,18]]`:

1. **Already sorted by start time**

2. **Initialize result with first interval:**
   - result = `[[1,3]]`

3. **Process [2,6]:**
   - Last interval: `[1,3]`
   - 2 ≤ 3 (overlaps)
   - Merge: `[1, max(3,6)] = [1,6]`
   - result = `[[1,6]]`

4. **Process [8,10]:**
   - Last interval: `[1,6]`
   - 8 > 6 (no overlap)
   - Add to result
   - result = `[[1,6],[8,10]]`

5. **Process [15,18]:**
   - Last interval: `[8,10]`
   - 15 > 10 (no overlap)
   - Add to result
   - result = `[[1,6],[8,10],[15,18]]`

### Why This is Optimal

- **Time Complexity O(n log n)**: Dominated by the sorting step. The merging pass is O(n)
- **Space Complexity O(n)**: We need to store the result, which in the worst case (no overlapping intervals) contains all n intervals. Sorting might use O(log n) stack space
- This is optimal because we must examine every interval at least once (O(n) lower bound), and comparison-based sorting cannot be faster than O(n log n)

### Edge Cases

1. **Single interval**: Return as-is
2. **No overlapping intervals**: Return all intervals (sorted)
3. **All intervals overlap**: Return one merged interval
4. **Intervals with same start time**: Sort by end time doesn't matter, merging handles it
5. **Nested intervals**: When one interval is completely inside another (e.g., [1,4] and [2,3])
6. **Touching intervals**: Intervals where end of one equals start of next (considered overlapping)
