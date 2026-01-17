import pytest
from product_of_array_except_self import product_except_self


def test_example_1():
    """Test case from example 1."""
    nums = [1, 2, 3, 4]
    result = product_except_self(nums)
    assert result == [24, 12, 8, 6]


def test_example_2():
    """Test case from example 2."""
    nums = [-1, 1, 0, -3, 3]
    result = product_except_self(nums)
    assert result == [0, 0, 9, 0, 0]


def test_two_elements():
    """Test with minimum input size."""
    nums = [2, 3]
    result = product_except_self(nums)
    assert result == [3, 2]


def test_all_ones():
    """Test with all elements being 1."""
    nums = [1, 1, 1, 1]
    result = product_except_self(nums)
    assert result == [1, 1, 1, 1]


def test_with_negative_numbers():
    """Test with negative numbers."""
    nums = [-2, -3, 4, -5]
    result = product_except_self(nums)
    assert result == [-60, -40, 30, -24]


def test_zero_at_beginning():
    """Test with zero at the beginning."""
    nums = [0, 2, 3, 4]
    result = product_except_self(nums)
    assert result == [24, 0, 0, 0]


def test_zero_at_end():
    """Test with zero at the end."""
    nums = [2, 3, 4, 0]
    result = product_except_self(nums)
    assert result == [0, 0, 0, 24]


def test_multiple_zeros():
    """Test with multiple zeros."""
    nums = [0, 0, 1, 2]
    result = product_except_self(nums)
    assert result == [0, 0, 0, 0]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
