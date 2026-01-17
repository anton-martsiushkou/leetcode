import pytest
from two_sum import two_sum


def test_example_1():
    """Test case from example 1."""
    nums = [2, 7, 11, 15]
    target = 9
    result = two_sum(nums, target)
    assert result == [0, 1]


def test_example_2():
    """Test case from example 2."""
    nums = [3, 2, 4]
    target = 6
    result = two_sum(nums, target)
    assert result == [1, 2]


def test_example_3():
    """Test case from example 3."""
    nums = [3, 3]
    target = 6
    result = two_sum(nums, target)
    assert result == [0, 1]


def test_negative_numbers():
    """Test with negative numbers."""
    nums = [-1, -2, -3, -4, -5]
    target = -8
    result = two_sum(nums, target)
    assert result == [2, 4]


def test_zero_in_array():
    """Test with zero in the array."""
    nums = [0, 4, 3, 0]
    target = 0
    result = two_sum(nums, target)
    assert result == [0, 3]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
