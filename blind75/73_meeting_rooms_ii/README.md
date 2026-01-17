# Meeting Rooms II

## Problem Description

Given an array of meeting time intervals `intervals` where `intervals[i] = [start_i, end_i]`, return the minimum number of conference rooms required.

## Examples

**Example 1:**
```
Input: intervals = [[0,30],[5,10],[15,20]]
Output: 2
Explanation: We need 2 rooms:
- Room 1: [0,30]
- Room 2: [5,10] and [15,20]
```

**Example 2:**
```
Input: intervals = [[7,10],[2,4]]
Output: 1
Explanation: Only one room is needed since the meetings don't overlap.
```

## Constraints

- `1 <= intervals.length <= 10^4`
- `0 <= start_i < end_i <= 10^6`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Min Heap | Sorted Array + Min Heap | O(n log n) | O(n) |
| Chronological Ordering | Two Sorted Arrays | O(n log n) | O(n) |

### Approach 1: Min Heap (Optimal)

Use a min heap to track when rooms become available. The heap stores the end times of ongoing meetings.

**Key Insight:** At any point, the number of rooms needed equals the number of overlapping meetings. We can track this by maintaining end times of all ongoing meetings.

### Algorithm Steps

1. Sort intervals by start time
2. Create a min heap to store end times of ongoing meetings
3. For each meeting:
   - If the earliest ending meeting (heap top) ends before or when current meeting starts, remove it (room becomes available)
   - Add the current meeting's end time to the heap
   - Track the maximum heap size (maximum concurrent meetings)
4. Return the maximum heap size

### Example Walkthrough

For `intervals = [[0,30],[5,10],[15,20]]`:

1. **Sort by start time**:
   - `[[0,30], [5,10], [15,20]]` (already sorted)

2. **Process meetings**:
   - Meeting [0,30]: heap = [30], max_rooms = 1
   - Meeting [5,10]: 5 < 30 (overlap), heap = [10, 30], max_rooms = 2
   - Meeting [15,20]: 15 >= 10 (remove 10), heap = [20, 30], max_rooms = 2

3. **Result**: 2 rooms needed

### Approach 2: Chronological Ordering

Separate start and end times, sort them, then process chronologically.

**Key Insight:** When a meeting starts, we need a room. When a meeting ends, a room becomes free. The maximum concurrent meetings is the answer.

### Algorithm Steps

1. Create separate arrays for start times and end times
2. Sort both arrays
3. Use two pointers to process events chronologically
4. When processing a start time: increment room count
5. When processing an end time: decrement room count
6. Track the maximum room count

### Why This is Optimal

- **Time Complexity O(n log n)**: Dominated by sorting
- **Space Complexity O(n)**: For the heap or separate arrays
- This is optimal because we must examine all meetings to determine the maximum overlap

### Alternative Approaches (Not Optimal)

1. **Brute Force**: Check every pair of meetings - O(n²) time
2. **Interval Tree**: More complex with similar time complexity
