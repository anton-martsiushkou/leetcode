from typing import List


def merge(intervals: List[List[int]]) -> List[List[int]]:
    """
    Merges all overlapping intervals.

    Time: O(n log n), Space: O(n)

    Args:
        intervals: List of intervals where each interval is [start, end]

    Returns:
        List of merged non-overlapping intervals
    """
    if not intervals:
        return intervals

    # Sort intervals by start time
    intervals.sort(key=lambda x: x[0])

    result = [intervals[0]]

    for i in range(1, len(intervals)):
        current = intervals[i]
        last = result[-1]

        # Check if current interval overlaps with last merged interval
        if current[0] <= last[1]:
            # Merge: extend the end time of last interval
            last[1] = max(last[1], current[1])
        else:
            # No overlap: add current interval
            result.append(current)

    return result
