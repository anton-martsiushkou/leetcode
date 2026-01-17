import pytest
from non_overlapping_intervals import erase_overlap_intervals


def test_example_1():
    """Test case from example 1."""
    intervals = [[1, 2], [2, 3], [3, 4], [1, 3]]
    result = erase_overlap_intervals(intervals)
    assert result == 1


def test_example_2_all_identical():
    """Test case from example 2 - all identical intervals."""
    intervals = [[1, 2], [1, 2], [1, 2]]
    result = erase_overlap_intervals(intervals)
    assert result == 2


def test_example_3_no_overlaps():
    """Test case from example 3 - no overlaps."""
    intervals = [[1, 2], [2, 3]]
    result = erase_overlap_intervals(intervals)
    assert result == 0


def test_example_4():
    """Test case from example 4."""
    intervals = [[1, 100], [11, 22], [1, 11], [2, 12]]
    result = erase_overlap_intervals(intervals)
    assert result == 2


def test_single_interval():
    """Test with a single interval."""
    intervals = [[1, 2]]
    result = erase_overlap_intervals(intervals)
    assert result == 0


def test_two_non_overlapping_intervals():
    """Test with two non-overlapping intervals."""
    intervals = [[1, 2], [3, 4]]
    result = erase_overlap_intervals(intervals)
    assert result == 0


def test_two_overlapping_intervals():
    """Test with two overlapping intervals."""
    intervals = [[1, 3], [2, 4]]
    result = erase_overlap_intervals(intervals)
    assert result == 1


def test_all_intervals_overlap_nested():
    """Test where all intervals overlap - nested case."""
    intervals = [[1, 10], [2, 9], [3, 8], [4, 7]]
    result = erase_overlap_intervals(intervals)
    assert result == 3


def test_multiple_groups_of_overlaps():
    """Test with multiple groups of overlapping intervals."""
    intervals = [[1, 2], [1, 3], [2, 3], [5, 6], [5, 7], [6, 7]]
    result = erase_overlap_intervals(intervals)
    assert result == 3


def test_touching_intervals_not_overlapping():
    """Test touching intervals which are not considered overlapping."""
    intervals = [[1, 2], [2, 3], [3, 4], [4, 5]]
    result = erase_overlap_intervals(intervals)
    assert result == 0


def test_reverse_sorted_intervals():
    """Test with reverse sorted intervals."""
    intervals = [[5, 6], [4, 5], [3, 4], [2, 3], [1, 2]]
    result = erase_overlap_intervals(intervals)
    assert result == 0


def test_complex_overlapping_case():
    """Test complex overlapping case."""
    intervals = [[0, 2], [1, 3], [2, 4], [3, 5], [4, 6]]
    result = erase_overlap_intervals(intervals)
    assert result == 2


def test_one_interval_contains_all_others():
    """Test where one large interval contains all others."""
    intervals = [[1, 100], [2, 3], [4, 5], [6, 7]]
    result = erase_overlap_intervals(intervals)
    assert result == 1


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
