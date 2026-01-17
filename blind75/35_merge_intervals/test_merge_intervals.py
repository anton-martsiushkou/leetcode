import pytest
from merge_intervals import merge


def test_example_1():
    """Test case from example 1."""
    intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]
    result = merge(intervals)
    assert result == [[1, 6], [8, 10], [15, 18]]


def test_example_2_touching_intervals():
    """Test case from example 2 - touching intervals."""
    intervals = [[1, 4], [4, 5]]
    result = merge(intervals)
    assert result == [[1, 5]]


def test_example_3_unsorted_intervals():
    """Test case from example 3 - unsorted intervals."""
    intervals = [[1, 4], [0, 4]]
    result = merge(intervals)
    assert result == [[0, 4]]


def test_example_4_nested_interval():
    """Test case from example 4 - nested interval."""
    intervals = [[1, 4], [2, 3]]
    result = merge(intervals)
    assert result == [[1, 4]]


def test_single_interval():
    """Test with a single interval."""
    intervals = [[1, 5]]
    result = merge(intervals)
    assert result == [[1, 5]]


def test_no_overlapping_intervals():
    """Test with no overlapping intervals."""
    intervals = [[1, 2], [3, 4], [5, 6]]
    result = merge(intervals)
    assert result == [[1, 2], [3, 4], [5, 6]]


def test_all_intervals_merge_into_one():
    """Test where all intervals merge into one."""
    intervals = [[1, 4], [2, 5], [3, 6]]
    result = merge(intervals)
    assert result == [[1, 6]]


def test_multiple_merges():
    """Test with multiple separate merges."""
    intervals = [[1, 3], [2, 4], [5, 7], [6, 8]]
    result = merge(intervals)
    assert result == [[1, 4], [5, 8]]


def test_completely_nested_intervals():
    """Test with completely nested intervals."""
    intervals = [[1, 10], [2, 3], [4, 5], [6, 7]]
    result = merge(intervals)
    assert result == [[1, 10]]


def test_unsorted_with_multiple_overlaps():
    """Test unsorted intervals with multiple overlaps."""
    intervals = [[2, 3], [4, 5], [6, 7], [8, 9], [1, 10]]
    result = merge(intervals)
    assert result == [[1, 10]]


def test_same_start_different_end():
    """Test intervals with same start but different end."""
    intervals = [[1, 4], [1, 5]]
    result = merge(intervals)
    assert result == [[1, 5]]


def test_empty_intervals():
    """Test with empty intervals list."""
    intervals = []
    result = merge(intervals)
    assert result == []


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
