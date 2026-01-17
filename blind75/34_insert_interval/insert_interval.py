from typing import List


def insert(intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
    """
    Inserts a new interval into a sorted list of non-overlapping intervals.
    Merges overlapping intervals if necessary.

    Time: O(n), Space: O(n)

    Args:
        intervals: List of non-overlapping intervals sorted by start time
        newInterval: New interval to insert

    Returns:
        List of intervals after insertion and merging
    """
    result = []
    i = 0
    n = len(intervals)

    # Phase 1: Add all intervals that come before newInterval
    while i < n and intervals[i][1] < newInterval[0]:
        result.append(intervals[i])
        i += 1

    # Phase 2: Merge all overlapping intervals with newInterval
    while i < n and intervals[i][0] <= newInterval[1]:
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i += 1
    result.append(newInterval)

    # Phase 3: Add all intervals that come after newInterval
    while i < n:
        result.append(intervals[i])
        i += 1

    return result
