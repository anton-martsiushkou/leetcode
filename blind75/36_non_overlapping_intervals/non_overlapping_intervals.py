from typing import List


def erase_overlap_intervals(intervals: List[List[int]]) -> int:
    """
    Returns the minimum number of intervals to remove to make
    the rest non-overlapping.

    Time: O(n log n), Space: O(1)

    Args:
        intervals: List of intervals where each interval is [start, end]

    Returns:
        Minimum number of intervals to remove
    """
    if not intervals:
        return 0

    # Sort intervals by end time
    intervals.sort(key=lambda x: x[1])

    count = 1  # We can always keep the first interval
    end = intervals[0][1]

    for i in range(1, len(intervals)):
        # If current interval doesn't overlap with the last kept interval
        if intervals[i][0] >= end:
            count += 1
            end = intervals[i][1]
        # Else: current interval overlaps, so we skip it (remove it)

    # Number of removals = total intervals - intervals we kept
    return len(intervals) - count
