import pytest
from maximum_product_subarray import max_product


def test_example_1():
    """Test case from example 1."""
    nums = [2, 3, -2, 4]
    result = max_product(nums)
    assert result == 6


def test_example_2():
    """Test case from example 2."""
    nums = [-2, 0, -1]
    result = max_product(nums)
    assert result == 0


def test_example_3():
    """Test case from example 3."""
    nums = [-2, 3, -4]
    result = max_product(nums)
    assert result == 24


def test_single_element_positive():
    """Test with single positive element."""
    nums = [5]
    result = max_product(nums)
    assert result == 5


def test_single_element_negative():
    """Test with single negative element."""
    nums = [-3]
    result = max_product(nums)
    assert result == -3


def test_all_positive():
    """Test with all positive numbers."""
    nums = [2, 3, 4]
    result = max_product(nums)
    assert result == 24


def test_all_negative_even_count():
    """Test with all negative numbers (even count)."""
    nums = [-2, -3, -4, -5]
    result = max_product(nums)
    assert result == 120


def test_all_negative_odd_count():
    """Test with all negative numbers (odd count)."""
    nums = [-2, -3, -4]
    result = max_product(nums)
    assert result == 12


def test_with_zeros():
    """Test with zeros in the array."""
    nums = [0, 2, 3]
    result = max_product(nums)
    assert result == 6


def test_negative_at_beginning():
    """Test with negative number at beginning."""
    nums = [-2, 3, 4]
    result = max_product(nums)
    assert result == 12


def test_multiple_zeros():
    """Test with multiple zeros."""
    nums = [0, 2, 0, 3]
    result = max_product(nums)
    assert result == 3


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
