import pytest
from maximum_subarray import max_sub_array


def test_example_1():
    """Test case from example 1."""
    nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    result = max_sub_array(nums)
    assert result == 6


def test_example_2():
    """Test case from example 2."""
    nums = [1]
    result = max_sub_array(nums)
    assert result == 1


def test_example_3():
    """Test case from example 3."""
    nums = [5, 4, -1, 7, 8]
    result = max_sub_array(nums)
    assert result == 23


def test_all_negative():
    """Test with all negative numbers."""
    nums = [-3, -2, -5, -1, -4]
    result = max_sub_array(nums)
    assert result == -1


def test_all_positive():
    """Test with all positive numbers."""
    nums = [1, 2, 3, 4, 5]
    result = max_sub_array(nums)
    assert result == 15


def test_single_negative():
    """Test with single negative number."""
    nums = [-1]
    result = max_sub_array(nums)
    assert result == -1


def test_alternating_signs():
    """Test with alternating positive and negative numbers."""
    nums = [-2, 3, -1, 4, -3, 2]
    result = max_sub_array(nums)
    assert result == 6


def test_zeros_and_negatives():
    """Test with zeros and negative numbers."""
    nums = [-2, 0, -1, 0, -3]
    result = max_sub_array(nums)
    assert result == 0


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
