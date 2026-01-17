from typing import List


def can_attend_meetings(intervals: List[List[int]]) -> bool:
    """
    Determines if a person can attend all meetings.
    Time: O(n log n), Space: O(1)

    Args:
        intervals: List of meeting time intervals [start, end]

    Returns:
        True if person can attend all meetings, False otherwise
    """
    if len(intervals) <= 1:
        return True

    # Sort intervals by start time
    intervals.sort(key=lambda x: x[0])

    # Check for overlaps in consecutive intervals
    for i in range(1, len(intervals)):
        if intervals[i][0] < intervals[i - 1][1]:
            return False  # Overlap found

    return True
