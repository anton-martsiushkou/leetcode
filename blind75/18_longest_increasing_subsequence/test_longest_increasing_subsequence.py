import pytest
from longest_increasing_subsequence import length_of_lis


def test_example_1():
    """Test case from example 1."""
    nums = [10, 9, 2, 5, 3, 7, 101, 18]
    result = length_of_lis(nums)
    assert result == 4


def test_example_2():
    """Test case from example 2."""
    nums = [0, 1, 0, 3, 2, 3]
    result = length_of_lis(nums)
    assert result == 4


def test_example_3():
    """Test case from example 3."""
    nums = [7, 7, 7, 7, 7, 7, 7]
    result = length_of_lis(nums)
    assert result == 1


def test_single_element():
    """Test with single element."""
    nums = [1]
    result = length_of_lis(nums)
    assert result == 1


def test_strictly_decreasing():
    """Test with strictly decreasing sequence."""
    nums = [5, 4, 3, 2, 1]
    result = length_of_lis(nums)
    assert result == 1


def test_strictly_increasing():
    """Test with strictly increasing sequence."""
    nums = [1, 2, 3, 4, 5]
    result = length_of_lis(nums)
    assert result == 5


def test_alternating():
    """Test with alternating values."""
    nums = [1, 3, 2, 4, 3, 5]
    result = length_of_lis(nums)
    assert result == 4


def test_negative_numbers():
    """Test with negative numbers."""
    nums = [-10, -5, 0, 3, 1, 2]
    result = length_of_lis(nums)
    assert result == 4


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
