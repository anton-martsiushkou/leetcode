import pytest
from meeting_rooms_ii import min_meeting_rooms, min_meeting_rooms_chronological


def test_example_1():
    """Test case from example 1."""
    intervals = [[0, 30], [5, 10], [15, 20]]
    assert min_meeting_rooms(intervals) == 2
    assert min_meeting_rooms_chronological(intervals) == 2


def test_example_2():
    """Test case from example 2."""
    intervals = [[7, 10], [2, 4]]
    assert min_meeting_rooms(intervals) == 1
    assert min_meeting_rooms_chronological(intervals) == 1


def test_empty_intervals():
    """Test with empty intervals."""
    intervals = []
    assert min_meeting_rooms(intervals) == 0
    assert min_meeting_rooms_chronological(intervals) == 0


def test_single_interval():
    """Test with single interval."""
    intervals = [[1, 5]]
    assert min_meeting_rooms(intervals) == 1
    assert min_meeting_rooms_chronological(intervals) == 1


def test_all_overlapping():
    """Test when all meetings overlap."""
    intervals = [[1, 10], [2, 9], [3, 8], [4, 7]]
    assert min_meeting_rooms(intervals) == 4
    assert min_meeting_rooms_chronological(intervals) == 4


def test_no_overlapping():
    """Test when no meetings overlap."""
    intervals = [[1, 5], [6, 10], [11, 15]]
    assert min_meeting_rooms(intervals) == 1
    assert min_meeting_rooms_chronological(intervals) == 1


def test_adjacent_meetings():
    """Test adjacent meetings that don't overlap."""
    intervals = [[1, 5], [5, 10], [10, 15]]
    assert min_meeting_rooms(intervals) == 1
    assert min_meeting_rooms_chronological(intervals) == 1


def test_complex_scenario():
    """Test complex overlapping scenario."""
    intervals = [[0, 30], [5, 10], [15, 20], [25, 35], [30, 40]]
    assert min_meeting_rooms(intervals) == 3
    assert min_meeting_rooms_chronological(intervals) == 3


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
