import pytest
from find_minimum_in_rotated_sorted_array import find_min


def test_example_1():
    """Test case from example 1."""
    nums = [3, 4, 5, 1, 2]
    result = find_min(nums)
    assert result == 1


def test_example_2():
    """Test case from example 2."""
    nums = [4, 5, 6, 7, 0, 1, 2]
    result = find_min(nums)
    assert result == 0


def test_example_3():
    """Test case from example 3."""
    nums = [11, 13, 15, 17]
    result = find_min(nums)
    assert result == 11


def test_single_element():
    """Test with single element."""
    nums = [1]
    result = find_min(nums)
    assert result == 1


def test_two_elements_rotated():
    """Test with two elements rotated."""
    nums = [2, 1]
    result = find_min(nums)
    assert result == 1


def test_two_elements_not_rotated():
    """Test with two elements not rotated."""
    nums = [1, 2]
    result = find_min(nums)
    assert result == 1


def test_rotated_once():
    """Test array rotated once."""
    nums = [2, 3, 4, 5, 1]
    result = find_min(nums)
    assert result == 1


def test_not_rotated():
    """Test array not rotated."""
    nums = [1, 2, 3, 4, 5]
    result = find_min(nums)
    assert result == 1


def test_negative_numbers():
    """Test with negative numbers."""
    nums = [3, 4, 5, -2, -1, 0, 1, 2]
    result = find_min(nums)
    assert result == -2


def test_large_rotation():
    """Test with large rotation."""
    nums = [5, 1, 2, 3, 4]
    result = find_min(nums)
    assert result == 1


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
