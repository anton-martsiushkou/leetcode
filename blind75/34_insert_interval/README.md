# Insert Interval

## Problem Description

You are given an array of non-overlapping intervals `intervals` where `intervals[i] = [starti, endi]` represent the start and the end of the `ith` interval and `intervals` is sorted in ascending order by `starti`. You are also given an interval `newInterval = [start, end]` that represents the start and end of another interval.

Insert `newInterval` into `intervals` such that `intervals` is still sorted in ascending order by `starti` and `intervals` still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return `intervals` after the insertion.

## Examples

**Example 1:**
```
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
```

**Example 2:**
```
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
```

**Example 3:**
```
Input: intervals = [], newInterval = [5,7]
Output: [[5,7]]
```

**Example 4:**
```
Input: intervals = [[1,5]], newInterval = [2,3]
Output: [[1,5]]
```

**Example 5:**
```
Input: intervals = [[1,5]], newInterval = [2,7]
Output: [[1,7]]
```

## Constraints

- `0 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 10^5`
- `intervals` is sorted by `starti` in ascending order
- `newInterval.length == 2`
- `0 <= start <= end <= 10^5`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Linear Scan with Merge | Array | O(n) | O(n) |

### Approach: Linear Scan with Merge

The solution processes the intervals in three phases: adding all intervals that come before the new interval, merging overlapping intervals with the new interval, and adding all intervals that come after.

**Key Insight:** Since the input intervals are already sorted and non-overlapping, we can process them in a single pass. We only need to identify which intervals overlap with the new interval and merge them.

### Algorithm Steps

1. Initialize an empty result array
2. **Phase 1 - Add intervals before new interval:**
   - Iterate through intervals where `interval.end < newInterval.start`
   - Add these intervals directly to the result (they don't overlap)
3. **Phase 2 - Merge overlapping intervals:**
   - While `interval.start <= newInterval.end` (they overlap):
     - Update `newInterval.start = min(interval.start, newInterval.start)`
     - Update `newInterval.end = max(interval.end, newInterval.end)`
   - Add the merged interval to the result
4. **Phase 3 - Add remaining intervals:**
   - Add all remaining intervals to the result (they come after the new interval)
5. Return the result

### Example Walkthrough

For `intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]]` and `newInterval = [4,8]`:

1. **Phase 1 - Before new interval:**
   - `[1,2]` ends at 2, which is < 4, so add to result: `[[1,2]]`

2. **Phase 2 - Merge overlapping:**
   - `[3,5]`: starts at 3 ≤ 8 (overlaps)
     - Merge: `newInterval = [min(3,4), max(5,8)] = [3,8]`
   - `[6,7]`: starts at 6 ≤ 8 (overlaps)
     - Merge: `newInterval = [min(6,3), max(7,8)] = [3,8]`
   - `[8,10]`: starts at 8 ≤ 8 (overlaps)
     - Merge: `newInterval = [min(8,3), max(10,8)] = [3,10]`
   - Add merged interval: `[[1,2],[3,10]]`

3. **Phase 3 - After new interval:**
   - `[12,16]` starts at 12 > 10, so add to result: `[[1,2],[3,10],[12,16]]`

### Why This is Optimal

- **Time Complexity O(n)**: We traverse the intervals array exactly once, examining each interval constant time
- **Space Complexity O(n)**: We need to store the result array, which in the worst case contains all original intervals plus the new one
- This is optimal because we must examine every interval at least once to determine if it overlaps with the new interval

### Edge Cases

1. **Empty intervals array**: Simply return the new interval
2. **New interval completely before all intervals**: Add it at the beginning
3. **New interval completely after all intervals**: Add it at the end
4. **New interval completely contained within an existing interval**: The existing interval remains unchanged
5. **New interval overlaps with all intervals**: All intervals merge into one
