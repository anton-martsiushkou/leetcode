# Meeting Rooms

## Problem Description

Given an array of meeting time intervals where `intervals[i] = [start_i, end_i]`, determine if a person could attend all meetings.

## Examples

**Example 1:**
```
Input: intervals = [[0,30],[5,10],[15,20]]
Output: false
Explanation: The person cannot attend all meetings because [0,30] overlaps with both [5,10] and [15,20].
```

**Example 2:**
```
Input: intervals = [[7,10],[2,4]]
Output: true
Explanation: The meetings don't overlap, so the person can attend all of them.
```

## Constraints

- `0 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= start_i < end_i <= 10^6`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Sort | Array | O(n log n) | O(1) or O(n) |

### Approach: Sort by Start Time

The optimal solution is to sort the intervals by start time, then check if any consecutive intervals overlap.

**Key Insight:** If we sort by start time, we only need to check if each meeting ends before the next one starts. If any meeting ends after the next one starts, there's an overlap.

### Algorithm Steps

1. Sort the intervals by start time
2. Iterate through the sorted intervals
3. For each interval (except the first), check if its start time is less than the previous interval's end time
4. If yes, there's an overlap - return false
5. If we finish the loop without finding overlaps, return true

### Example Walkthrough

For `intervals = [[0,30],[5,10],[15,20]]`:

1. **Sort by start time**:
   - `[[0,30], [5,10], [15,20]]` (already sorted)

2. **Check overlaps**:
   - Compare [0,30] and [5,10]: 5 < 30 → Overlap! Return false

For `intervals = [[7,10],[2,4]]`:

1. **Sort by start time**:
   - `[[2,4], [7,10]]`

2. **Check overlaps**:
   - Compare [2,4] and [7,10]: 7 >= 4 → No overlap
   - Return true

### Why This is Optimal

- **Time Complexity O(n log n)**: Dominated by sorting. The checking phase is O(n).
- **Space Complexity O(1)**: If we can sort in-place, otherwise O(n) for the sorting algorithm
- This is optimal because we must examine all intervals to determine if there are any conflicts

### Edge Cases

1. **Empty intervals**: Return true (no meetings to attend)
2. **Single interval**: Return true (only one meeting)
3. **Adjacent meetings**: `[[1,5],[5,10]]` should return true (meetings touch but don't overlap)
