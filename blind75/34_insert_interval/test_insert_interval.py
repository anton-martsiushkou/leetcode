import pytest
from insert_interval import insert


def test_example_1():
    """Test case from example 1."""
    intervals = [[1, 3], [6, 9]]
    newInterval = [2, 5]
    result = insert(intervals, newInterval)
    assert result == [[1, 5], [6, 9]]


def test_example_2():
    """Test case from example 2."""
    intervals = [[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]]
    newInterval = [4, 8]
    result = insert(intervals, newInterval)
    assert result == [[1, 2], [3, 10], [12, 16]]


def test_example_3_empty_intervals():
    """Test case from example 3 - empty intervals."""
    intervals = []
    newInterval = [5, 7]
    result = insert(intervals, newInterval)
    assert result == [[5, 7]]


def test_example_4_new_interval_inside_existing():
    """Test case from example 4 - new interval inside existing."""
    intervals = [[1, 5]]
    newInterval = [2, 3]
    result = insert(intervals, newInterval)
    assert result == [[1, 5]]


def test_example_5_new_interval_extends_existing():
    """Test case from example 5 - new interval extends existing."""
    intervals = [[1, 5]]
    newInterval = [2, 7]
    result = insert(intervals, newInterval)
    assert result == [[1, 7]]


def test_new_interval_before_all():
    """Test with new interval before all existing intervals."""
    intervals = [[3, 5], [6, 9]]
    newInterval = [1, 2]
    result = insert(intervals, newInterval)
    assert result == [[1, 2], [3, 5], [6, 9]]


def test_new_interval_after_all():
    """Test with new interval after all existing intervals."""
    intervals = [[1, 2], [3, 5]]
    newInterval = [6, 9]
    result = insert(intervals, newInterval)
    assert result == [[1, 2], [3, 5], [6, 9]]


def test_merge_all_intervals():
    """Test merging all intervals into one."""
    intervals = [[1, 2], [3, 4], [5, 6]]
    newInterval = [0, 7]
    result = insert(intervals, newInterval)
    assert result == [[0, 7]]


def test_adjacent_intervals_no_overlap():
    """Test adjacent intervals with no overlap."""
    intervals = [[1, 5]]
    newInterval = [6, 8]
    result = insert(intervals, newInterval)
    assert result == [[1, 5], [6, 8]]


def test_adjacent_intervals_touching():
    """Test adjacent intervals that are touching."""
    intervals = [[1, 5]]
    newInterval = [5, 8]
    result = insert(intervals, newInterval)
    assert result == [[1, 8]]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
