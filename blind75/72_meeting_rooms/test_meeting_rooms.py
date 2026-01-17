import pytest
from meeting_rooms import can_attend_meetings


def test_example_1_overlapping_meetings():
    """Test case from example 1 - overlapping meetings."""
    intervals = [[0, 30], [5, 10], [15, 20]]
    assert can_attend_meetings(intervals) == False


def test_example_2_non_overlapping_meetings():
    """Test case from example 2 - non-overlapping meetings."""
    intervals = [[7, 10], [2, 4]]
    assert can_attend_meetings(intervals) == True


def test_empty_intervals():
    """Test with empty intervals."""
    intervals = []
    assert can_attend_meetings(intervals) == True


def test_single_interval():
    """Test with single interval."""
    intervals = [[1, 5]]
    assert can_attend_meetings(intervals) == True


def test_adjacent_meetings_no_overlap():
    """Test adjacent meetings that don't overlap."""
    intervals = [[1, 5], [5, 10]]
    assert can_attend_meetings(intervals) == True


def test_adjacent_meetings_with_gap():
    """Test adjacent meetings with gaps."""
    intervals = [[1, 5], [6, 10], [11, 15]]
    assert can_attend_meetings(intervals) == True


def test_multiple_overlaps():
    """Test with multiple overlapping meetings."""
    intervals = [[1, 10], [2, 6], [3, 5], [7, 9]]
    assert can_attend_meetings(intervals) == False


def test_meetings_in_reverse_order():
    """Test with meetings provided in reverse chronological order."""
    intervals = [[15, 20], [10, 15], [5, 10]]
    assert can_attend_meetings(intervals) == True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
