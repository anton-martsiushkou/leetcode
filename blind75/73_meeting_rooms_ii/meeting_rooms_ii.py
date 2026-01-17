import heapq
from typing import List


def min_meeting_rooms(intervals: List[List[int]]) -> int:
    """
    Returns the minimum number of conference rooms required.
    Time: O(n log n), Space: O(n)

    Args:
        intervals: List of meeting time intervals [start, end]

    Returns:
        Minimum number of conference rooms needed
    """
    if not intervals:
        return 0

    # Sort intervals by start time
    intervals.sort(key=lambda x: x[0])

    # Min heap to track end times of ongoing meetings
    heap = []
    max_rooms = 0

    for interval in intervals:
        # Remove meetings that have ended
        while heap and heap[0] <= interval[0]:
            heapq.heappop(heap)

        # Add current meeting's end time
        heapq.heappush(heap, interval[1])

        # Track maximum concurrent meetings
        max_rooms = max(max_rooms, len(heap))

    return max_rooms


def min_meeting_rooms_chronological(intervals: List[List[int]]) -> int:
    """
    Returns the minimum number of conference rooms using chronological ordering.
    Time: O(n log n), Space: O(n)

    Args:
        intervals: List of meeting time intervals [start, end]

    Returns:
        Minimum number of conference rooms needed
    """
    if not intervals:
        return 0

    starts = sorted([interval[0] for interval in intervals])
    ends = sorted([interval[1] for interval in intervals])

    rooms = 0
    max_rooms = 0
    start_ptr = 0
    end_ptr = 0

    while start_ptr < len(starts):
        if starts[start_ptr] < ends[end_ptr]:
            # Meeting starts, need a room
            rooms += 1
            start_ptr += 1
            max_rooms = max(max_rooms, rooms)
        else:
            # Meeting ends, free a room
            rooms -= 1
            end_ptr += 1

    return max_rooms
