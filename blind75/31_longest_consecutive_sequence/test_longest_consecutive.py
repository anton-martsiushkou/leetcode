import pytest
from longest_consecutive import longest_consecutive, longest_consecutive_sorting


def test_example_1():
    """Test case from example 1."""
    nums = [100, 4, 200, 1, 3, 2]
    assert longest_consecutive(nums) == 4


def test_example_2():
    """Test case from example 2."""
    nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
    assert longest_consecutive(nums) == 9


def test_empty_array():
    """Test with empty array."""
    nums = []
    assert longest_consecutive(nums) == 0


def test_single_element():
    """Test with single element."""
    nums = [1]
    assert longest_consecutive(nums) == 1


def test_all_duplicates():
    """Test with all duplicates."""
    nums = [1, 1, 1, 1]
    assert longest_consecutive(nums) == 1


def test_no_consecutive():
    """Test with no consecutive numbers."""
    nums = [1, 3, 5, 7, 9]
    assert longest_consecutive(nums) == 1


def test_negative_numbers():
    """Test with negative numbers."""
    nums = [-1, -2, 0, 1, 2]
    assert longest_consecutive(nums) == 5


def test_consecutive_at_end():
    """Test with consecutive sequence at end."""
    nums = [10, 5, 1, 2, 3, 4]
    assert longest_consecutive(nums) == 4


def test_sorting_example_1():
    """Test sorting implementation with example 1."""
    nums = [100, 4, 200, 1, 3, 2]
    assert longest_consecutive_sorting(nums) == 4


def test_sorting_example_2():
    """Test sorting implementation with example 2."""
    nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
    assert longest_consecutive_sorting(nums) == 9


def test_sorting_empty():
    """Test sorting with empty array."""
    nums = []
    assert longest_consecutive_sorting(nums) == 0


def test_sorting_duplicates():
    """Test sorting with duplicates."""
    nums = [1, 1, 1, 1]
    assert longest_consecutive_sorting(nums) == 1


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
